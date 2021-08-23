resource "aws_iam_server_certificate" "iam_server_certificates_cert" {
  name = "iam_server_certificates_${var.test_prefix}${var.test_suffix}"

  certificate_body = <<EOF
-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUV81BN2CYVWrXFZPJRu8E/DBp+PIwDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMTA4MTgxMzU2MzJaFw0zMTA4
MTYxMzU2MzJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQD8yHhqJIhdnNpUXam6Iyo4iHo9KIO2fz4ZCYZFBJ5h
oSOXSa3klNavqbDI581PhSV538KEMfdxOLRqogWlNGK2nCBsnmn68ASMQVduEEI3
bKFbq2tsKBX+fNe3gf1gqRUXRygjyqZbneBFNDT4raEZ2IziZH0iJSYoiwX67vNO
l3vebK2kRCKYOWKe8Y5lg9mVs55zvC4ZzZZEczYmVsvxnCJ7mP5+S8cgJTJyUaKJ
qQly42NNWiDmHO+clsGdAwyLj7f5a7X8oixdnTnGkpyhxEyyTEwHtBD025WuBE5g
wIE/43ZFW7CAJcoaoqnRUsBR9/UxE6T8DdqiyxoYtHKrAgMBAAGjUzBRMB0GA1Ud
DgQWBBTjZ9Q0q50HC4E+i2hMjhNv1O01XDAfBgNVHSMEGDAWgBTjZ9Q0q50HC4E+
i2hMjhNv1O01XDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQDD
IQgLqmweLNYMtMOH/P1rv8VXffAibG4SSuLw6XBtamNpBJUsvsOjUxjaXufFxX8+
7eT8b7W6TBoV2arzTQIHGBtJ/e5hPcpCFOEtZldNUhCV91LXBBaauF4H7AZlpmEL
fOEXkQS+w26zhgB034NiIOJQne3h6++qZPTyT/Abm4/FnKKxq8mEZRK7zvQEjL5g
9UyXFDO4jOQBxCrQWiz1bdHfaCfTgYFf7cxjf49zYW9mxYqxek7lwnb2uTqmeIak
2Wl6snwgm2R9ZJlSn+6XMJMzuqdC5hXUC7bAF+pwYj1eMSTs6dGElpR46JU1kRGI
YoLcOjm4Um6uFcSZ+lcG
-----END CERTIFICATE-----
EOF

  private_key = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA/Mh4aiSIXZzaVF2puiMqOIh6PSiDtn8+GQmGRQSeYaEjl0mt
5JTWr6mwyOfNT4Uled/ChDH3cTi0aqIFpTRitpwgbJ5p+vAEjEFXbhBCN2yhW6tr
bCgV/nzXt4H9YKkVF0coI8qmW53gRTQ0+K2hGdiM4mR9IiUmKIsF+u7zTpd73myt
pEQimDlinvGOZYPZlbOec7wuGc2WRHM2JlbL8Zwie5j+fkvHICUyclGiiakJcuNj
TVog5hzvnJbBnQMMi4+3+Wu1/KIsXZ05xpKcocRMskxMB7QQ9NuVrgROYMCBP+N2
RVuwgCXKGqKp0VLAUff1MROk/A3aossaGLRyqwIDAQABAoIBAH6q1ewaMlr2ZlEw
PgSP6nA5s0dCLf4c9LkHFMkw70xsurDssr5/9rQ/i96giTj3tzfC8G3du+h4Sa/F
UJ7gCTcINRc4qlKcQysk8vKsIwIy5QhSdZnU3HCRjmbeNGFjWLY5advFiGcQTXlg
F8sk9j1FThOD11Bs3Rojz/NWGYWfU/iB5a3o38Ry8jFvMm+RqgvjIFmSQW9gsysu
rF5OnQOjxphy6ryDyshN3ActEgVuHBQtcz21GJ0Awae5ygO3X6FQWJZ53Dm09Im5
AKIXqMruYUNKm85RTd06mitkq0NtYKzxcHDatfjGB71Fk/9yVAIpj4PqjpWmwfUF
lIki3AECgYEA/v/n8w9n9HHnsg+CK51qgZSeV+/I64lSFgBMcFMjR4cLsuUW2+Tw
jt43u+PIJmwkQaH2v9FTiElSRWfnhxxZp7qePWKvP/PNiM5MuvuoH2WTqMshrdQ6
r8HmO+qgZtYduuTDJERjGBzpkeWzhmemvXrWopkNQCTQyRkfKU+f+MECgYEA/cZW
mCib9uq9xnpYd+6rcpcGlk2slE877u4jeJSwt1logZPwjh4Divsz9jylWguujU1D
o7WVi4/vKFzBzJN8B6P60Ygvv8g2n/P8ySrj6yks/ZFjIDh8VyuImCRPtZaKy6OD
X4TIJfyd8DbTcHRJY2fUykm+SGbRnsKszS+r+msCgYEAjqso4BG/jMFp63LH1pUl
MMiw6uxKkpVq6spR+gpZ07wX7IhGCNOHT7e+oQ3pPq4EfwAaKOn5/Wdgc1qV7D8N
Xk2IGTosaPCED0W1ImyIfMB9I9Q6zKHQD+PBZ8Z25fJoAdfZ4mGsu6H5gFpu3gUC
AkNG8QQGqNuwQMMgbyGQlEECgYEA1P4Lneevh1zb7OKwr/Bca2/AcpQ8vbmGxvGO
Sd0aOLjM0ry1EMK2HIbbxNZ/vYTCewwXMibsndcvVgQMEETUu6Dwmb4gvWkwF63z
Q3BBHRVREn1c4iUmjw8VtQP89p6kXWgekiPizKcqc/vAGMkw80KpQSMFlsTdOM+R
syFIyYMCgYA2GDA8cZgNpRzvEfmZiC5nKb+bx0jnF675Hd/f5Woe5ik15lN7PMFs
LkFOoMQlxHHOUKwnKM5IS9NqUmHL+Aq4q4Id4s89UK6ta2g2+DnebX0+rbxvBoFo
64rLunnaT9tWJP+oc/q0sWKJnNGP/XKnTZ9RV5/MuI7+aTN1x6Acxg==
-----END RSA PRIVATE KEY-----
EOF
}