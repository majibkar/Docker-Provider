/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file is an exact copy of
// plugin/pkg/auth/authorizer/webhook/certs_test.go

package webhook

var caCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDCzCCAfOgAwIBAgIJAKK9m2Cfg5uhMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNV
BAMMEHdlYmhvb2tfYXV0aHpfY2EwIBcNMTYwMjE2MjM0NDI4WhgPMjI4OTEyMDEy
MzQ0MjhaMBsxGTAXBgNVBAMMEHdlYmhvb2tfYXV0aHpfY2EwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDohVcY9fnI/ZDpMCVBchCxq/NKpn3veJm+lAEZ
Mjoz3yBNJWZ4FgTejtLx7pZMjcMKdbCpzwx3WtP3YW1pAvC/64w1eKMuKjlFUGqN
QRu6cPkif/5P/LmwT0/cLNxbLoISEi85EBmPpSW+BnNtHAj3X4RUVHmN3wlTR78v
0OzWMtkXJDApzUAOj5OP07iXRa2lGWuHNOCiivgLd8BBmbSar6WVOfEKbD6WHLy3
zxs8bVwzB10O1RNul61z77YlnrZDcXWepxfURwcOJ1HTHeP7ffnsEuRdErgxCxYt
gvN5ndfWCFZdPTrYfzro9whX4xfgfnyykw3lxViTlhOd4aKFAgMBAAGjUDBOMB0G
A1UdDgQWBBSumZL6MMwmFGyhQAwl/v0lYDzdZjAfBgNVHSMEGDAWgBSumZL6MMwm
FGyhQAwl/v0lYDzdZjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAG
6k+bZxKYq4PVZHWTKA7RSjv95FMMr4RSFwKn/n8TUD44ANWYqDrEfVmxAMn3NVK9
ckA8mIRym4IGiWD9eBGgPNNtbAq8Wl/9+5qbDMerpXuRnG3wNY7RU75Rl008m52r
c2i86ZPUi2fAJZyMf5StWE21oKiDYYQqlB6xxsIj6OHhf7536vEysoztNX5FpS2n
q8wG0EhJVhG+Qyww8IlZA5Cjoh71Eqkcwb4cuLjPypxmLm0ywZ/6KgzV+IF+CT2v
TJIpMokDUKlRi9cWSqkWXFE6xbCmhrrwKYsi0X6Vvi7a0pmOnSzKCQl8jN8u4A9R
xar2YeJ6mCCzSAPM69DP
-----END CERTIFICATE-----`)

var badCACert = []byte(`-----BEGIN CERTIFICATE-----
MIIDCzCCAfOgAwIBAgIJAPqJyUfmRxGLMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNV
BAMMEHdlYmhvb2tfYXV0aHpfY2EwIBcNMTYwMjE2MjM0NDI4WhgPMjI4OTEyMDEy
MzQ0MjhaMBsxGTAXBgNVBAMMEHdlYmhvb2tfYXV0aHpfY2EwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQCift1FXgGvXR0tx/zOCz+E4ei4Qu94Wvy25ZUg
WK/FJ2dpzaTp0ziJDa5hL2dNfE2HIkn9Nf8wcRhziBJeipRu3H2MTdUDiYIoOE7t
wm0JjGrRBX8+VAKIA+OfmdQy30AJ2KLv+5MUOXjpsr0rpqO8od8ehZ5homwXniTE
GrkgyP5BoHadz6ltrTRsYbT2xhwfIPftAf66DHrjH7bXExiPYsH+4+ipkzSSPUYc
/ECzwaoWrY+DCS/nz0CUeq3aIwJMCoQPaNyDgV5LavUN5gfAbIYA99ZN0DxFn6Bn
8/QaP3VP7gH2ytHsHJtQWN/UrERy89cG8/mvw7A6JiNLyHH/AgMBAAGjUDBOMB0G
A1UdDgQWBBS6IGeGHZCylibt0GzY0dP6C0J9VjAfBgNVHSMEGDAWgBS6IGeGHZCy
libt0GzY0dP6C0J9VjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQAi
A1dp75kbePFZsUNjxN6B/Pv0vSoaOjQkc4hpxKbI4VRCuPGmMRFYTlKCzoZ53OqQ
2Jmu1Zbzel/bV5vXrW0BOfUpfWYzd/usIJEuTgU8ijBIB+IHAXYwwxeKRcz3C+7+
9RBMF7gSg9pU2hrSvjhh7Q96IMJ42Z7tI3WD8SZaQLjY1NW1jrQVsg66ktdMke7x
zC8oIRIBH4W6l5s7jtZx1k305NE04pigcFLxCxOmicKd66ysI5hAZkD7y0dgwgtL
IqCQy6t7uJDydRiNRfPFr9Eg7uOu83JGw11f3bGVhJVCbzHyKddvkQsQbdaMHRgZ
zgmWLORg+ls1H1oaJiNW
-----END CERTIFICATE-----`)

var serverKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAtegsP499au5ZxlwM26rk3TnRgakchQi/9bhfMr0LaEKng1lR
XopzzGuGeZQswzbx7iiH89JzFkurZoEmZwtS4Aybit92VOSv0EUnyx7WR3V21ObZ
iQO0rr0UmG84NjdzATkqF+R5Z+HN9shwgBI4PR1j/ybCt7jNz+OM/VmqsgzoKLoa
bGrx7LCTPk8y5G8AoPOrIAP+9WHJsKQSRT8Lru4lYqseBxvhjqo8NRqzZLg79ldY
aKFqa2N2zr5qp94sG3/zihNDxjZvyyn9c8qvPBL0xOyayvOJG8eZUmjQpUMv7Jk3
qFmdMgGaDJRw0Qg6+/Zt6MHNs6Rbb8hmwuMSpwIDAQABAoIBAQCjzeFijwzKKL4w
0B1IBhi3WeReFPG4nkt1ssQPBYrrJPKBZgHO13A1STI78wFn/OdYpajfF8hI8HT1
BiGVsu27Eb9TC60b/x6OtmeCEk+044LRbtu+9NZUb7HHHogI0l++X0KXZ0coE38L
1izwNvfrmLa+QaIgHMtAg9EnJwJ993n4L31GovWh8MGmVyJX/F92y+agNwWkNYYp
iLWFyon+HbNVL13WOOYnYEdA8Me3+Gucy1EOfWMF7mgmuO2vcfnxXd6b16VjAwtE
jGCQfzgpWGHLpgwoBgDmnPUbdNPUT3MbA9jqG2mlnBSBQveYgKrmFdDYnAjnCM4L
uF2ztBzhAoGBAOYc3sF3YjpIIMsyH9omqtfOuxO+oZkpb2vB9kgdXCDcG870M+BC
bNzV7DCSV8QAUqjKQK1r3gq62UZMLXZbG8x5UnM8/EK0X1CSqygwSWjGpYxIQEhh
O2lq69WipkNDnX1ZmrvEdHD2cxqkkXZ7bdRKRasrFJgvJa3XbiJ18KYxAoGBAMpe
/72EcX9oL3KT8tJSpvasrw17p/XkMMCxTp3IDb3krF/4k5bYF61F68/LNSy3xkos
ZrPUK/U160iuHSYCpMq4pPmlWgKq4hmUMOt+8Yy622zDlugarq9VLqvSdGHm+r6F
5fHilXB0UsTXXOuLZWLcSQ0MBgiaVCLb2AmXZhhXAoGAEjSchw/r7JKCTbE0hezj
PVm0wVYmsNhvYUYiNwhjnpHrfU8iv45h0IL4QcuCOBaSc5o0zcOn+I9Z207xldiV
dXLvzAA6MQjWNai08+QGGs0EkfmxZEiVC70S1X8dylqSHjW1oT9kuv80khoNDCOt
x8rsgiNRaMzqHTvbEczk8jECgYB2Od+wSULBSw2FI5fVdcHjFGlEODycs44j1LH4
DZqxmHl3q9IVavMSIGouQCo1kLuAM8ZgQpDXtYNaN5YB0cOSRyLiUc5vBoQGq4OU
4Nme/L8aIH315TiuZ9ZXPSEO3REZ40G9+UCSrPJ52tOHLC2z/ruSqraPqhGDN+pT
WCamCwKBgEPa+kVrPs0khQH8+sbFbU9ifj4fhPAiSwj2fKuXFro2mE205vAMHye/
SYs/mPzYzKSd7F+7Zk6oVrgFVskTiReW3phF+cIl+CdcnIenF0jW1PVgGw8znu+P
SbHSdqV+tB7AW2J7sH8TZtfMUPAK2MJ4S+1uaHK86K79ym4Rz0E2
-----END RSA PRIVATE KEY-----`)

var serverCert = []byte(`-----BEGIN CERTIFICATE-----
MIIC/zCCAeegAwIBAgIJAN7rkfhaX8FZMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNV
BAMMEHdlYmhvb2tfYXV0aHpfY2EwIBcNMTYwMjE2MjM0NDI4WhgPMjI4OTEyMDEy
MzQ0MjhaMB8xHTAbBgNVBAMMFHdlYmhvb2tfYXV0aHpfc2VydmVyMIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtegsP499au5ZxlwM26rk3TnRgakchQi/
9bhfMr0LaEKng1lRXopzzGuGeZQswzbx7iiH89JzFkurZoEmZwtS4Aybit92VOSv
0EUnyx7WR3V21ObZiQO0rr0UmG84NjdzATkqF+R5Z+HN9shwgBI4PR1j/ybCt7jN
z+OM/VmqsgzoKLoabGrx7LCTPk8y5G8AoPOrIAP+9WHJsKQSRT8Lru4lYqseBxvh
jqo8NRqzZLg79ldYaKFqa2N2zr5qp94sG3/zihNDxjZvyyn9c8qvPBL0xOyayvOJ
G8eZUmjQpUMv7Jk3qFmdMgGaDJRw0Qg6+/Zt6MHNs6Rbb8hmwuMSpwIDAQABo0Aw
PjAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DATBgNVHSUEDDAKBggrBgEFBQcDATAP
BgNVHREECDAGhwR/AAABMA0GCSqGSIb3DQEBCwUAA4IBAQCZHB9UCl2CfylWP3db
xUamawnRoTYlsOcUh4f2tlHMY+vYiEStN+LECk62YpeaHl/nz/lk7g1Jx9aua39z
wFIHiXYhwSWOtgmzpbxYLye1yajKXbbA1T7mEZJTjewDB9i1LcB9W3EV5VJ8Y1GY
AYKuKQ4Cb1HrqLsrw/1PDm0VouWzf2ESv8CBvAv/pYLVfwgS6WsUqn9wycpLEnqQ
RK66/AoiOaxUIjEP0O1q6pi6Mag7XAfeNtx8J0VGt4cRG4rvWCbKVUyvKfUCkipN
gJu09S+KIz3x1CJLRuJX9tB+cFnnykDLQ2IKg7x44O83ikNk8+Di3iT/awCguWPE
rHh5
-----END CERTIFICATE-----`)

var clientKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA5ij4WXWGvbmfAYhEafKRvLEHSkUCYIDjwQAlnHoLf/lz+Fh2
DEv4lcBaycwk3+LVUGKgYOg91txYJvGD3HcmVThXZvcgJd4V9Ll3aY/6xVRCenWi
UNgVQVQITGkMn09ZkSXbZCK4wqz9oTVh0Ti5a7apOS2V07yL0q7vw003v5TBqzC/
FgRwE0bv1rKYYQ80WbDlYkkYGf216zQTwS4g/nShCZAX9eqSfbBg6B/A3OwpbIfx
09BWuwWhp5QnS4w002gGWavRFNzu8pUHUv6zMN8OKpasv+Na+ZB+gMt4+e2Y7qNz
76QL23eGwc6oWn8lQBtkDLmLIa6jbWX067U76QIDAQABAoIBAQCJpGzJSzC2W8DM
sMqBNdCUMKZ0cwq13b7W2BimGJKyCOOi3HxUZEaYf/2Leyt+PPBm72SML7dzvDh3
qa269gKVqmkSqa2vF763qQbRuYo14msTQzA7+s3TUMbZs2UaDOE6nZIzs1QdEElp
1DvYXHz+/rD7Adj9VF+mMnouqQoy5kgJTnVZ8sOyl/9R6F67xKBIvcrtPfqVZzuG
2hGAMUnawxFUajQC7BynIeCWrk79SUmQgilyNgRdY6+rGh2uRupIxuiAukPtuag1
Li+wnNl1UGECtv9ZnnboKvg2334k5vhYScGRJbwbr7Zt3ZaNd0Z/DE9kTtnhBS7v
9qWdc7CBAoGBAPR4hz1fhHFiPmMEAGuiNms6WdyIfyonIRYas8ZDKUQGdxn/aO8a
CURktHRlm6iYT+j1cbf3RnLEN9pNr3V2EySOMc+rXUNifcP7Vl53akAQmISUfQWG
UfwaNLicbavf6m9UCiwWByAZghqDZSLiwmLHIjGcSJQiFuhZryioDydxAoGBAPED
q1Z7oNhzwRYie9OB5ylnrCH8G3yFl8egBmQrPJKIQHA9mAGg01LEJwQNoWewyAWx
jfeFtWvIgZkj49cluZgHYyF81jApaNraxtXAgIwC1n7oAIttmeklZ/V1HntknG3Y
ow2bV/NA3aPOTPYxW8oDv7U9lvwve7kIFxeWjE/5AoGASfXI3G1wUSkqvKPySJ3b
ntcZZpm49xS9csWDS+D3tAfMsoXNxkB3O0TIP0qaLAhgbJcM314k5wWr7BSCl6Ow
KOgH887hOUirycXZHF0+PMGIktulcy1u0jlPZ+aTW2MztpiTN0E2yKRO8xx7VXGK
431hP+cLIh2qFoNDdaZaZ1ECgYEArw++PWQxMefqgVxs2vXJZY7TPiA0Ct+ynqKC
4fFx3vGu9JgYuF4MAVtPB6eq7HlA4LnWZ8ssOuz6DbU/AoB5bY84FxPpNDRv4D/3
Gz3nYUuSZ72234+tsuaju2vlxzUOVs97qB+E48Di/N+VkWHKzVKpxkjFScpnsL/K
niyRIGkCgYEAriuxbOCczL/j6u2Xq1ngEsGg+RXjtOYGoJWo7B8qlVL4nF8w1Nbd
FxEmOChQgUnBdwb93qHCSq0Fidf7OfewrfJJkstWIh3zPS4umLZo7R3YblncpdfT
M197uckIWccZml2jF/c7nvK+MjwDRhkOl2a6HzMxcdBwYUJmSwmIZ4k=
-----END RSA PRIVATE KEY-----`)

var clientCert = []byte(`-----BEGIN CERTIFICATE-----
MIIC7jCCAdagAwIBAgIJAN7rkfhaX8FaMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNV
BAMMEHdlYmhvb2tfYXV0aHpfY2EwIBcNMTYwMjE2MjM0NDI4WhgPMjI4OTEyMDEy
MzQ0MjhaMB8xHTAbBgNVBAMMFHdlYmhvb2tfYXV0aHpfY2xpZW50MIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5ij4WXWGvbmfAYhEafKRvLEHSkUCYIDj
wQAlnHoLf/lz+Fh2DEv4lcBaycwk3+LVUGKgYOg91txYJvGD3HcmVThXZvcgJd4V
9Ll3aY/6xVRCenWiUNgVQVQITGkMn09ZkSXbZCK4wqz9oTVh0Ti5a7apOS2V07yL
0q7vw003v5TBqzC/FgRwE0bv1rKYYQ80WbDlYkkYGf216zQTwS4g/nShCZAX9eqS
fbBg6B/A3OwpbIfx09BWuwWhp5QnS4w002gGWavRFNzu8pUHUv6zMN8OKpasv+Na
+ZB+gMt4+e2Y7qNz76QL23eGwc6oWn8lQBtkDLmLIa6jbWX067U76QIDAQABoy8w
LTAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DATBgNVHSUEDDAKBggrBgEFBQcDAjAN
BgkqhkiG9w0BAQsFAAOCAQEA2IZNhkVrSTAIeP2N2WzOHqbFbGyO+NA8G9Hb5fiX
e1YS2Ku3ERYNr+HLxNHCsXiSUKjjBmXMc4z0XaHJznEKEbotZftjTlTQlHi3/5vm
dIG18pmO/E5ebVXl6pU96v/hBd8N5rWp9WUKgP0y59r/JA+oNpmd10A+RyaOyrFK
rBm8Z8rvDYMrXSpOwx9BNDuhqzbdG8MYw5vO55Er3hwTXoapsMqSh5s9+OFFpUJi
2uEoQlwWiYRtQj6g4wgr4woDEbv8XxsHqGfs+GSnmRsB69xRI24lEtC+nS6Rz3Sh
YWeN0gD8PsQC1KJVv6xCGo1yXSEwytRMB23XYtAZahLdLg==
-----END CERTIFICATE-----`)
