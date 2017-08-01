database-setup: need_root
	echo 'UPDATE user SET plugin="";' | mysql mysql | true
	echo 'CREATE DATABASE $(DB_NAME);' | mysql mysql | true
	echo 'CREATE USER "$(DB_USER)"@localhost;' | mysql | true
	echo 'GRANT ALL PRIVILEGES ON * . * TO "$(DB_USER)"@localhost;' | mysql | true
	echo 'UPDATE user SET password=PASSWORD("$(DB_PASS)") WHERE user="$(DB_USER)";' | mysql mysql | true
	echo 'FLUSH PRIVILEGES;' | mysql | true

broker-setup: need_root
	[ "1" -eq "$(shell rabbitmq-plugins enable rabbitmq_management | grep 'Plugin configuration unchanged' | wc -l)" ] || (rabbitmqctl stop_app && rabbitmqctl start_app)
	rabbitmqctl add_user $(R_USER) $(R_PASS) || rabbitmqctl change_password $(R_USER) $(R_PASS)
	rabbitmqctl set_user_tags $(R_USER) administrator
	rabbitmqctl set_permissions -p / $(R_USER) ".*" ".*" ".*"
	wget -O /usr/bin/rabbitmqadmin http://127.0.0.1:15672/cli/rabbitmqadmin
	chmod a+x /usr/bin/rabbitmqadmin
	rabbitmqadmin declare queue name=dlx-queue
	rabbitmqadmin declare exchange name=dlx-exchange type=topic
	rabbitmqctl set_policy DLX ".*" '{"dead-letter-exchange":"dlx-exchange"}' --apply-to queues
	rabbitmqadmin declare binding source="dlx-exchange" destination_type="queue" destination="dlx-queue" routing_key="#"
