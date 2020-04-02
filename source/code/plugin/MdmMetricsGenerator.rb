#!/usr/local/bin/ruby
# frozen_string_literal: true

class MdmMetricsGenerator
  require "logger"
  require "json"
  require_relative "MdmAlertTemplates"
  require_relative "ApplicationInsightsUtility"
  require_relative "constants"

  @log_path = "/var/opt/microsoft/docker-cimprov/log/mdm_metrics_generator.log"
  @log = Logger.new(@log_path, 1, 5000000)

  # @@oom_killed_container_count_metric_name = "OomKilledContainerCount"
  # @@container_restart_count_metric_name = "ContainerRestartCount"
  @oom_killed_container_count_hash = {}
  @container_restart_count_hash = {}
  @pod_ready_hash = {}
  @pod_not_ready_hash = {}
  @pod_ready_percentage_hash = {}

  # @@cpu_usage_milli_cores = "cpuUsageMillicores"
  # @@cpu_usage_nano_cores = "cpuUsageNanoCores"

  @@metric_name_metric_percentage_name_hash = {
    Constants::CPU_USAGE_MILLI_CORES => "cpuUsagePercentage",
    Constants::CPU_USAGE_NANO_CORES => "cpuUsagePercentage",
    "memoryRssBytes" => "memoryRssPercentage",
    "memoryWorkingSetBytes" => "memoryWorkingSetPercentage",
  }

  def initialize
  end

  class << self
    # def appendPodMetrics(records, batch_time)
    #   begin
    #     @log.info "in appendPodMetrics..."
    #     @log.info "oom killed container count: #{@oom_killed_container_count_hash.length}"
    #     if !@oom_killed_container_count_hash.empty?
    #       @oom_killed_container_count_hash.each { |key, value|
    #         key_elements = key.split("~~")
    #         if key_elements.length != 2
    #           next
    #         end

    #         # get dimension values by key
    #         podControllerNameDimValue = key_elements[0]
    #         podNamespaceDimValue = key_elements[1]

    #         record = MdmAlertTemplates::Oom_killed_container_count_custom_metrics_template % {
    #           timestamp: batch_time,
    #           metricName: MdmMetrics::OOM_KILLED_CONTAINER_COUNT,
    #           controllerNameDimValue: podControllerNameDimValue,
    #           namespaceDimValue: podNamespaceDimValue,
    #           containerCountMetricValue: value,
    #         }
    #         records.push(JSON.parse(record))
    #       }
    #     else
    #       @log.info "No OOMKilled containers found"
    #     end
    #   rescue => errorStr
    #     @log.info "Error appending pod metrics for metric: #{MdmMetrics::OOM_KILLED_CONTAINER_COUNT} : #{errorStr}"
    #     ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
    #   end
    #   @log.info "Done appending PodMetrics for oom killed containers..."
    #   @oom_killed_container_count_hash = {}
    #   return records
    # end
    def populatePodReadyPercentageHash
      begin
        @log.info "in populatePodReadyPercentageHash..."
        @pod_ready_hash.each { |dim_key, value|
          podsNotReady = @pod_not_ready_hash.key?(dim_key) ? @pod_not_ready_hash[dim_key] : 0
          totalPods = value + podsNotReady
          podsReadyPercentage = (value / totalPods) * 100
          @pod_ready_percentage_hash[dim_key] = podsReadyPercentage
          # Deleting this key value pair from not ready hash,
          # so that we can get those dimensions for which there are 100% of the pods in not ready state
          if (@pod_not_ready_hash.key?(dim_key))
            @pod_not_ready_hash.delete(dim_key)
          end
        }

        # Add 0% pod ready for these dimensions
        if @pod_not_ready_hash.length > 0
          @pod_not_ready_hash.each { |key, value|
            @pod_ready_percentage_hash[key] = 0
          }
        end

        # Cleaning up hashes after use
        @pod_ready_hash = {}
        @pod_not_ready_hash = {}
      rescue => errorStr
        @log.info "Error in populatePodReadyPercentageHash: #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
    end

    def appendPodMetrics(records, metricName, metricHash, batch_time)
      begin
        @log.info "in appendPodMetrics..."
        if !metricHash.empty?
          metricHash.each { |key, value|
            key_elements = key.split("~~")
            if key_elements.length != 2
              next
            end

            # get dimension values by key
            podControllerNameDimValue = key_elements[0]
            podNamespaceDimValue = key_elements[1]

            record = MdmAlertTemplates::Pod_Metrics_custom_metrics_template % {
              timestamp: batch_time,
              metricName: metricName,
              controllerNameDimValue: podControllerNameDimValue,
              namespaceDimValue: podNamespaceDimValue,
              containerCountMetricValue: value,
            }
            records.push(JSON.parse(record))
          }
        else
          @log.info "No records found in hash for metric: #{metricName}"
        end
      rescue => errorStr
        @log.info "Error appending pod metrics for metric: #{metricName} : #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
      @log.info "Done appending PodMetrics for metric: #{metricName}..."
      return records
    end

    def appendAllPodMetrics(records, batch_time)
      begin
        @log.info "in appendAllPodMetrics..."
        @log.info "@oom_killed_container_count_hash: #{@oom_killed_container_count_hash}"
        records = appendPodMetrics(records, Constants::MDM_OOM_KILLED_CONTAINER_COUNT, @oom_killed_container_count_hash, batch_time)
        @oom_killed_container_count_hash = {}
        @log.info "@container_restart_count_hash: #{@container_restart_count_hash}"
        records = appendPodMetrics(records, Constants::MDM_CONTAINER_RESTART_COUNT, @container_restart_count_hash, batch_time)
        @container_restart_count_hash = {}
        # Computer the percentage here, because we need to do this after all chunks have been processed.
        populatePodReadyPercentageHash
        @log.info "@pod_ready_percentage_hash: #{@pod_ready_percentage_hash}"
        records = appendPodMetrics(records, Constants::MDM_POD_READY_PERCENTAGE, @pod_ready_percentage_hash, batch_time)
        @pod_ready_percentage_hash = {}
      rescue => errorStr
        @log.info "Error in appendAllPodMetrics: #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
      return records
    end

    def getContainerResourceUtilMetricRecords(record, metricName, percentageMetricValue, dims)
      records = []
      begin
        dimElements = dims.split("~~")
        if dimElements.length != 4
          return records
        end

        # get dimension values
        containerName = dimElements[0]
        podName = dimElements[1]
        controllerName = dimElements[2]
        podNamespace = dimElements[3]

        resourceUtilRecord = MdmAlertTemplates::Container_resource_utilization_template % {
          timestamp: record["DataItems"][0]["Timestamp"],
          metricName: @@metric_name_metric_percentage_name_hash[metricName],
          containerNameDimValue: containerName,
          podNameDimValue: podName,
          controllerNameDimValue: controllerName,
          namespaceDimValue: podNamespace,
          containerResourceUtilizationPercentage: percentageMetricValue,
        }
        records.push(JSON.parse(resourceUtilRecord))
      rescue => errorStr
        @log.info "Error in getContainerResourceUtilMetricRecords: #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
      return records
    end

    def getNodeResourceMetricRecords(record, metric_name, metric_value, percentage_metric_value)
      records = []
      begin
        custommetricrecord = MdmAlertTemplates::Node_resource_metrics_template % {
          timestamp: record["DataItems"][0]["Timestamp"],
          metricName: metric_name,
          hostvalue: record["DataItems"][0]["Host"],
          objectnamevalue: record["DataItems"][0]["ObjectName"],
          instancenamevalue: record["DataItems"][0]["InstanceName"],
          metricminvalue: metric_value,
          metricmaxvalue: metric_value,
          metricsumvalue: metric_value,
        }
        records.push(JSON.parse(custommetricrecord))

        if !percentage_metric_value.nil?
          additional_record = MdmAlertTemplates::Node_resource_metrics_template % {
            timestamp: record["DataItems"][0]["Timestamp"],
            metricName: @@metric_name_metric_percentage_name_hash[metric_name],
            hostvalue: record["DataItems"][0]["Host"],
            objectnamevalue: record["DataItems"][0]["ObjectName"],
            instancenamevalue: record["DataItems"][0]["InstanceName"],
            metricminvalue: percentage_metric_value,
            metricmaxvalue: percentage_metric_value,
            metricsumvalue: percentage_metric_value,
          }
          records.push(JSON.parse(additional_record))
        end
      rescue => errorStr
        @log.info "Error in getNodeResourceMetricRecords: #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
      return records
    end

    def generateOOMKilledPodMetrics(podControllerName, podNamespace)
      begin
        dim_key = [podControllerName, podNamespace].join("~~")
        @log.info "adding dimension key to oom killed container hash..."
        @oom_killed_container_count_hash[dim_key] = @oom_killed_container_count_hash.key?(dim_key) ? @oom_killed_container_count_hash[dim_key] + 1 : 1
      rescue => errorStr
        @log.warn "Error in generateOOMKilledPodMetrics: #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
    end

    def generateContainerRestartsMetrics(podControllerName, podNamespace)
      begin
        dim_key = [podControllerName, podNamespace].join("~~")
        @log.info "adding dimension key to container restart count hash..."
        @container_restart_count_hash[dim_key] = @container_restart_count_hash.key?(dim_key) ? @container_restart_count_hash[dim_key] + 1 : 1
      rescue => errorStr
        @log.warn "Error in generateContainerRestartsMetrics: #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
    end

    def generatePodReadyMetrics(podControllerName, podNamespace, podReadyCondition)
      begin
        dim_key = [podControllerName, podNamespace].join("~~")
        @log.info "adding dimension key to pod ready hash..."
        if podReadyCondition == true
          @pod_ready_hash[dim_key] = @pod_ready_hash.key?(dim_key) ? @pod_ready_hash[dim_key] + 1 : 1
        else
          @pod_not_ready_hash[dim_key] = @pod_not_ready_hash.key?(dim_key) ? @pod_not_ready_hash[dim_key] + 1 : 1
        end
      rescue => errorStr
        @log.warn "Error in generatePodReadyMetrics: #{errorStr}"
        ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
      end
    end

    def generateStaleJobCountMetrics(podControllerNameDimValue,
      podNamespaceDimValue)

    # def generatePodMetrics(metricName, podControllerName, podNamespace)
    #   begin
    #     @log.info "in generatePodMetrics..."
    #     # group by distinct dimension values
    #     dim_key = [podControllerName, podNamespace].join("~~")
    #     if metricName == Constants::MDM_OOM_KILLED_CONTAINER_COUNT
    #       @log.info "adding dimension key to oom killed container hash..."
    #       @oom_killed_container_count_hash[dim_key] = @oom_killed_container_count_hash.key?(dim_key) ? @oom_killed_container_count_hash[dim_key] + 1 : 1
    #     elsif metricName == Constants::MDM_CONTAINER_RESTART_COUNT
    #       @log.info "adding dimension key to container restart count hash..."
    #       @container_restart_count_hash[dim_key] = @container_restart_count_hash.key?(dim_key) ? @container_restart_count_hash[dim_key] + 1 : 1
    #       # if !@container_restart_count_hash.key?(dim_key) ?
    #       #   @container_restart_count_hash[dim_key] = metricValue
    #       # end
    #     end
    #   rescue => errorStr
    #     @log.warn "Error in generatePodMetrics: #{errorStr}"
    #     ApplicationInsightsUtility.sendExceptionTelemetry(errorStr)
    #   end
    # end
  end
end
