rabbitmqadmin declare exchange name=api type=topic -u admin -padmin
rabbitmqadmin declare queue --vhost="/" name=messages durable=true -u admin -padmin
rabbitmqadmin --vhost="/" declare binding source=api destination_type=queue destination=messages routing_key="api.conversation.*.message.*.added" -u admin -padmin
