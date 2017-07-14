#!/bin/bash
echo "Content-type: text/html"
echo ""
echo "<html>"
echo "<head>"
echo "<title>"
echo "Update Status Page"
echo "</title>"
echo "</head>"
echo "<body bgcolor="#cccccc" text="#000000">"

{ YOUR COMMAND HERE MAKE SURE IS SAVE, USE SUDO }  > /dev/null 2>&1
# example
# sudo /var/lib/cgi-bin/update_status

if (( $? == 0 )) ; then
    echo "update succesfully"
else
    echo "update failed"
fi
echo "<input type=\"button\" value=\"Back\" onclick=\"location.href = document.referrer; return false;\">"
echo "</body>"
echo "</html>"
exit 0

# example sudo file
# start
#	Cmnd_Alias  STATUS_PAGE = /var/lib/cgi-bin/update_status
#	www-data ALL=(ALL) NOPASSWD: STATUS_PAGE
# end

# example calling script /var/lib/cgi-bin/update_status
# start
#	/usr/local/sbin/status-page --config /etc/status-page.yaml > /var/lib/www/instances_status.html
# end

# example nginx config (need fcgiwrap installed)
# start
#	location ~ \.cgi$ {
#		root /var/lib/cgi-bin;
#		rewrite ^/cgi-bin/(.*)$ /$1;
#		auth_basic "Restricted Access";
#		auth_basic_user_file /etc/nginx/users.passwd;
#		include /etc/nginx/fastcgi_params;
#		fastcgi_param AUTH_USER $remote_user;
#		fastcgi_param REMOTE_USER $remote_user;
#		fastcgi_param SCRIPT_FILENAME $document_root/$fastcgi_script_name;
#		fastcgi_pass fcgiwrap;
#	}
# end

# and add the upstream config like this
# start
#	upstream fcgiwrap {
#		server unix:/var/run/fcgiwrap.socket;
#	}
# end
