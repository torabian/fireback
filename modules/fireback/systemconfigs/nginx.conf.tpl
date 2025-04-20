# We are generating the nginx config for you.
# We did not change any nginx definition for you, only printing this as
# a help for you to copy and paste in your conf file

# Usually located here:
# nano /etc/nginx/sites-enabled/default

# IMPORTANT: in order to Smart UI work correctly, make sure location is between two slashes:
# /authentication/ - first and last slashes are important


location {{ .Location}} {
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-Host $host;
    proxy_set_header X-Forwarded-Port $server_port;
    proxy_pass http://{{ .Host}}:{{ .Port}}/;
}