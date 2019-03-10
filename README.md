# UCNB - Unifi Controller Notification Bridge
UCNB is a helper application built to extend the Unifi Controller notifications by allowing more output methods other than eMail.
In it core this is just an SMTP server with some logic to forward eMails to other resources. This was initially built to forward notifications
to IFTTT so that they trigger notifications in a phone, but it is possible to use this project for many other purposes.

## Getting Started
These instructions are meant to get you up and running with your very own UCNB to the point where you can send a test notification with [IFTTT](http://ifttt.com).

### Installing
The easiest way to get started is to download the precompiled binaries found under releases. You will also need a IFTTT webhook setup if you haven't already, my suggestion is to have the webhook trigger a notification initially just to verify that everything is working.
UCNB requires a  minimal configuration to start the application to work with IFTTT webhooks. All that you need to change is the `event_name` and `api_key`.
Please consider using a better username and password combination, as these were only chosen for demonstration purposes.

```bash
ucnb -username=username -password=password -output-plugin='{"type": "ifttt", "value": {"event_name": "<event_name>", "api_key": "<api_key>"}}'
```

This will start a SMTP server bound to `127.0.0.1:1025` and forward any notification eMail it receives from the Unifi Controller to the IFTTT webhook.
To bind to a different address or port simply specify it when starting ucnf. The following example will listen to `0.0.0.0:25`.

```bash
ucnb -addr=0.0.0.0:25 -username=username -password=password -output-plugin='{"type": "ifttt", "value": {"event_name": "<event_name>", "api_key": "<api_key>"}}'

```
UCNB should now be running, make sure that it is listening to a port and interface that is accessible by the device running the cloud controller software. Sign into the cloud controller and go to the settings page. Under `Settings>Controller>Mail Server` check `Enable mail server` and fill in the ip, port, username and password that you have configured ucnb with. If you have the IFTTT webhook setup properly you should be able to send a test notification and see it trigger in IFTTT. The notification you should receive should be similar to the following message.

```
You have received a message generated by UniFi Controller Your SMTP Server settings appear to be working correctly!

Controller URL: https://unifi.yourdomain.com:8443/manage/site/default
```

Under `Settings>Notifications` you should be able to select which events that you want to receive notifications for. As UCNB works by forwarding email notifications, you will need to check the email box, to receive a IFTTT notification for that event.

### Docker
There are Docker images built to make deployment of UCNB event easier, and they are available for arm, arm64 and amd64 architectures. Go to [Docker Hub](https://hub.docker.com/r/phillebaba/ucnb) for more information about how to use the Docker image.

## Output Plugin
Currently there are only two output plugins that are supported. This is mainly because I created this project so that I could use it with IFTTT. The output plugins are written in a modular way so that it is possible to add more in the future. If you need another method feel free to contribute. The output plugin is configured with the `-output-plugin` option when running UCNB.

### IFTTT
This output method will trigger a IFTTT webhook with the specified `event_name` and authenticating with the `api_key"`. Currently the whole email is sent as one value to the webhook, `value1`, but there is work to add support for parsing the text and splitting it into two values.

```json
{
    "type": "ifttt",
    "value": {
        "event_name": "test_event",
        "api_key": "test_key"
    }
}
```

### HTTP
Will send a POST request to the specified `endpoint` with the notification message in the request body field `message`.
```json
{
    "type": "http",
    "value": {
        "endpoint": "https://example.com"
    }
}
```

The example above would create a POST request that can be represented as the following curl command.
```bash
curl -X POST --header "application/x-www-form-urlencoded" --data "message=<message_data>" https://example.com
```

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
