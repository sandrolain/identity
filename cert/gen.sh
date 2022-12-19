rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=IT/ST=Padua/L=Padua/O=Sandro Lain/OU=Development/CN=*.sandrolain.com/emailAddress=sandrolain@outlook.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text


# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout admin-server-key.pem -out admin-server-req.pem -subj "/C=IT/ST=Padua/L=Padua/O=Sandro Lain/OU=Development/CN=*.sandrolain.com/emailAddress=sandrolain@outlook.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in admin-server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out admin-server-cert.pem -extfile server-ext.cnf

echo "Admin Server's signed certificate"
openssl x509 -in admin-server-cert.pem -noout -text


# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-server-key.pem -out client-server-req.pem -subj "/C=IT/ST=Padua/L=Padua/O=Sandro Lain/OU=Development/CN=*.sandrolain.com/emailAddress=sandrolain@outlook.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in client-server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-server-cert.pem -extfile server-ext.cnf

echo "Client Server's signed certificate"
openssl x509 -in client-server-cert.pem -noout -text