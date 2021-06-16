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
DEPULSO_WRITE_LOCATION="/tmp/"
```