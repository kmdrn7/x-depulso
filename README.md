# Depulso
<hr>
Yet another log rotator but for ether traffics (generate PCAP)

## Environment Variables
This is an example of .env

```
DEPULSO_LISTEN_INTERFACE="eno1"
DEPULSO_INTERVAL=60
DEPULSO_CRON_SPEC="0 * * * * *"
DEPULSO_PROMISC=true
DEPULSO_DAYS_RETENTION=2
DEPULSO_WRITE_CSV_LOCATION="/data/csv/"
DEPULSO_KAFKA_TOPIC="topic_name"
DEPULSO_KAFKA_HOST="127.0.0.1"
DEPULSO_KAFKA_PORT="29092"
DEPULSO_WRITE_LOCATION="/tmp/"
DEPULSO_CICFLOWMETER_PATH="/app/CICFlowmeter.jar"
```