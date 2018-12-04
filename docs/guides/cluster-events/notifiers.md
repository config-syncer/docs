---
title: Notifiers
description: Notifiers
menu:
  product_kubed_0.9.0:
    identifier: events-notifier
    name: Notifiers
    parent: cluster-events
    weight: 15
product_name: kubed
menu_name: product_kubed_0.9.0
section_menu_id: guides
---

> New to Kubed? Please start [here](/docs/concepts/README.md).

# Supported Notifiers

Kubed can send notifications via Email, SMS or Chat for various operations using [appscode/go-notify](https://github.com/appscode/go-notify) library. To connect to these services, you need to create a Secret with the appropriate keys. Then pass the secret name to Kubed by setting `notifierSecretName` field in Kubed cluster config. Also, set `clusterName` to a meaningful name for you. This cluster name will be prefixed to any notification sent via Email/SMS/Chat so that you can identify the source easily.

## Hipchat
To receive chat notifications in Hipchat, create a Secret with the following key:

| Name                         | Description                                                   |
|------------------------------|---------------------------------------------------------------|
| HIPCHAT_AUTH_TOKEN           | `Required` Hipchat [api access token](https://developer.atlassian.com/hipchat/guide/hipchat-rest-api/api-access-tokens). You can use room notification tokens, if you are planning to send notifications to a single room.   |
| HIPCHAT_BASE_URL             | `Optional` Base url of Hipchat server                         |
| HIPCHAT_CA_CERT_DATA         | `Optional` PEM encoded CA certificate used by Hipchat server  |
| HIPCHAT_INSECURE_SKIP_VERIFY | `Optional` If set to `true`, skips SSL verification           |
| HIPCHAT_NOTIFY               | `Optional` If set to `true`, a notification will be triggered |
| HIPCHAT_COLOR                | `Optional` Set color of message sent. [Possible values](https://github.com/tbruyelle/hipchat-go/blob/35aebc99209aa41498ab4823a83e254037f062fd/hipchat/hipchat.go#L107): yellow, green, red, purple, gray, random. |

```console
$ echo -n 'your-hipchat-auth-token' > HIPCHAT_AUTH_TOKEN
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./HIPCHAT_AUTH_TOKEN
secret "notifier-config" created
```
```yaml
apiVersion: v1
data:
  HIPCHAT_AUTH_TOKEN: eW91ci1oaXBjaGF0LWF1dGgtdG9rZW4=
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T01:54:37Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "2244"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: 372bc159-70dc-11e7-9b0b-080027503732
type: Opaque
```

Now, to receive notifications via Hipchat, configure receiver as below:

 - notifier: `Hipchat`
 - to: a list of chat room names

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Hipchat
    to:
    - ops-alerts
notifierSecretName: notifier-config
```


## Stride
To receive chat notifications in Stride, create a Secret with the following key:

| Name                          | Description                                                  |
|-------------------------------|--------------------------------------------------------------|
| STRIDE_CLOUD_ID               | `Required` You can get stride cloud_id from your organization's URL   |
| STRIDE_ROOM_TOKEN             | `Optional` Stride [room token](https://developer.atlassian.com/cloud/stride/security/authentication/#using-room-tokens)   |

If you login into your stride web app, you will see an URL with this format `https://app.stride.com/{{cloud ID}}/lobby`. From here, you will get cloud_id.

```console
$ echo -n 'your-stride-cloud-id' > STRIDE_CLOUD_ID
$ echo -n 'your-stride-room-token' > STRIDE_ROOM_TOKEN
$ kubectl create secret generic notifier-config -n demo \
    --from-file=./STRIDE_CLOUD_ID \
    --from-file=./STRIDE_ROOM_TOKEN
secret "notifier-config" created
```

```yaml
apiVersion: v1
data:
  STRIDE_CLOUD_ID: eW91ci1zdHJpZGUtY2xvdWQtaWQ=
  STRIDE_ROOM_TOKEN: eW91ci1zdHJpZGUtcm9vbS10b2tlbg==
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T01:54:37Z
  name: notifier-config
  namespace: demo
  resourceVersion: "2244"
  selfLink: /api/v1/namespaces/demo/secrets/notifier-config
  uid: 372bc159-70dc-11e7-9b0b-080027503732
type: Opaque
```

Now, to receiver notifications via Stride, configure receiver as below:

- notifier: `Stride`
- to: a single room name

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Stride
    to:
    - ops-alerts
notifierSecretName: notifier-config
```

### Using Stride App

You can also use [Stride App](https://developer.atlassian.com/cloud/stride/security/authentication/#building-an-app) as client to send notification. This will allow you to send notifications to multiple rooms.

You need to do followings:

- [Create an App](https://developer.atlassian.com/apps/create) and get `{clientId}` and `{clientSecret}`.
- [Install this app](https://developer.atlassian.com/cloud/stride/getting-started/#install-the-app) in conversations.

To receive chat notifications in stride using App, create a Secret with the following key:

| Name                          | Description                                                           |
|-------------------------------|-----------------------------------------------------------------------|
| STRIDE_CLOUD_ID               | `Required` You can get stride cloud_id from your organization's URL   |
| STRIDE_CLIENT_ID              | `Optional` Stride app clientId                                        |
| STRIDE_CLIENT_SECRET          | `Optional` Stride app clientSecret                                    |


```console
$ echo -n 'your-stride-cloud-id' > STRIDE_CLOUD_ID
$ echo -n 'your-stride-app-client-id' > STRIDE_CLIENT_ID
$ echo -n 'your-stride-app-client-secret' > STRIDE_CLIENT_SECRET

$ kubectl create secret generic notifier-config -n demo \
    --from-file=./STRIDE_CLOUD_ID \
    --from-file=./STRIDE_CLIENT_ID \
    --from-file=./STRIDE_CLIENT_SECRET
secret "notifier-config" created
```

```yaml
apiVersion: v1
data:
  STRIDE_CLOUD_ID: eW91ci1zdHJpZGUtY2xvdWQtaWQ=
  STRIDE_CLIENT_ID: eW91ci1zdHJpZGUtYXBwLWNsaWVudC1pZA==
  STRIDE_CLIENT_SECRET: eW91ci1zdHJpZGUtYXBwLWNsaWVudC1zZWNyZXQ=
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T01:54:37Z
  name: notifier-config
  namespace: demo
  resourceVersion: "2244"
  selfLink: /api/v1/namespaces/demo/secrets/notifier-config
  uid: 372bc159-70dc-11e7-9b0b-080027503732
type: Opaque
```

Now, to receiver notifications via Stride App, configure receiver as below:

- notifier: `Stride`
- to: a list of room names where this app is installed

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Stride
    to:
    - ops-alerts
    - dev-alerts
notifierSecretName: notifier-config
```


## Mailgun
To receive email notifications via Mailgun, create a Secret with the following keys:

| Name                    | Description                                  |
|-------------------------|----------------------------------------------|
| MAILGUN_DOMAIN          | `Required` domain name for Mailgun account   |
| MAILGUN_API_KEY         | `Required` Mailgun API Key                   |
| MAILGUN_FROM            | `Required` sender email address              |
| MAILGUN_PUBLIC_API_KEY  | `Optional` Mailgun public API Key            |

```console
$ echo -n 'your-mailgun-domain' > MAILGUN_DOMAIN
$ echo -n 'no-reply@example.com' > MAILGUN_FROM
$ echo -n 'your-mailgun-api-key' > MAILGUN_API_KEY
$ echo -n 'your-mailgun-public-api-key' > MAILGUN_PUBLIC_API_KEY
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./MAILGUN_DOMAIN \
    --from-file=./MAILGUN_FROM \
    --from-file=./MAILGUN_API_KEY \
    --from-file=./MAILGUN_PUBLIC_API_KEY
secret "notifier-config" created
```
```yaml
apiVersion: v1
data:
  MAILGUN_API_KEY: eW91ci1tYWlsZ3VuLWFwaS1rZXk=
  MAILGUN_DOMAIN: eW91ci1tYWlsZ3VuLWRvbWFpbg==
  MAILGUN_FROM: bm8tcmVwbHlAZXhhbXBsZS5jb20=
  MAILGUN_PUBLIC_API_KEY: bWFpbGd1bi1wdWJsaWMtYXBpLWtleQ==
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T01:31:24Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "714"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: f8e91037-70d8-11e7-9b0b-080027503732
type: Opaque
```

Now, to receive notifications via Mailgun, configure receiver as below:
 - notifier: `Mailgun`
 - to: a list of email addresses

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Mailgun
    to:
    - ops-alerts@example.com
notifierSecretName: notifier-config
```


## SMTP
To configure any email provider, use SMTP notifier. To receive email notifications via SMTP services, create a Secret with the following keys:

| Name                      | Description                                           |
|---------------------------|-------------------------------------------------------|
| SMTP_HOST                 | `Required` Host address of smtp server                |
| SMTP_PORT                 | `Required` Port of smtp server                        |
| SMTP_INSECURE_SKIP_VERIFY | `Required` If set to `true`, skips SSL verification   |
| SMTP_USERNAME             | `Required` Username                                   |
| SMTP_PASSWORD             | `Required` Password                                   |
| SMTP_FROM                 | `Required` Sender email address                       |


```console
$ echo -n 'your-smtp-host' > SMTP_HOST
$ echo -n 'your-smtp-port' > SMTP_PORT
$ echo -n 'your-smtp-insecure-skip-verify' > SMTP_INSECURE_SKIP_VERIFY
$ echo -n 'your-smtp-username' > SMTP_USERNAME
$ echo -n 'your-smtp-password' > SMTP_PASSWORD
$ echo -n 'your-smtp-from' > SMTP_FROM
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./SMTP_HOST \
    --from-file=./SMTP_PORT \
    --from-file=./SMTP_INSECURE_SKIP_VERIFY \
    --from-file=./SMTP_USERNAME \
    --from-file=./SMTP_PASSWORD \
    --from-file=./SMTP_FROM
secret "notifier-config" created
```

To configure Kubed to send email notifications using a GMail account, set the Secrets like below:
```
$ echo -n 'smtp.gmail.com' > SMTP_HOST
$ echo -n '587' > SMTP_PORT
$ echo -n 'your-gmail-adress' > SMTP_USERNAME
$ echo -n 'your-gmail-password' > SMTP_PASSWORD
$ echo -n 'your-gmail-address' > SMTP_FROM
```

Now, to receive notifications via SMTP, configure receiver as below:

 - notifier: `SMTP`
 - to: a list of email addresses

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: SMTP
    to:
    - ops-alerts@example.com
notifierSecretName: notifier-config
```


## Twilio
To receive SMS notifications via Twilio, create a Secret with the following keys:

| Name                | Description                                        |
|---------------------|----------------------------------------------------|
| TWILIO_ACCOUNT_SID  | `Required` Twilio account SID                      |
| TWILIO_AUTH_TOKEN   | `Required` Twilio authentication token             |
| TWILIO_FROM         | `Required` Sender mobile number                    |

```console
$ echo -n 'your-twilio-account-sid' > TWILIO_ACCOUNT_SID
$ echo -n 'your-twilio-auth-token' > TWILIO_AUTH_TOKEN
$ echo -n 'your-twilio-from' > TWILIO_FROM
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./TWILIO_ACCOUNT_SID \
    --from-file=./TWILIO_AUTH_TOKEN \
    --from-file=./TWILIO_FROM
secret "notifier-config" created
```
```yaml
apiVersion: v1
data:
  TWILIO_ACCOUNT_SID: eW91ci10d2lsaW8tYWNjb3VudC1zaWQ=
  TWILIO_AUTH_TOKEN: eW91ci10d2lsaW8tYXV0aC10b2tlbg==
  TWILIO_FROM: eW91ci10d2lsaW8tZnJvbQ==
kind: Secret
metadata:
  creationTimestamp: 2017-07-26T17:38:38Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "27787"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: 41f57a61-7229-11e7-af79-08002738e55e
type: Opaque
```

Now, to receive notifications via Twilio, configure receiver as below:

 - notifier: `Twilio`
 - to: a list of receiver mobile numbers

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Twilio
    to:
    - +1-999-888-1234
notifierSecretName: notifier-config
```


## Slack
To receive chat notifications in Slack, create a Secret with the following keys:

| Name             | Description                      |
|------------------|----------------------------------|
| SLACK_AUTH_TOKEN | `Required` Slack [legacy auth token](https://api.slack.com/custom-integrations/legacy-tokens) |

```console
$ echo -n 'your-slack-auth-token' > SLACK_AUTH_TOKEN
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./SLACK_AUTH_TOKEN
secret "notifier-config" created
```
```yaml
apiVersion: v1
data:
  SLACK_AUTH_TOKEN: eW91ci1zbGFjay1hdXRoLXRva2Vu
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T01:58:58Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "2534"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: d2571817-70dc-11e7-9b0b-080027503732
type: Opaque
```

Now, to receive notifications via Slack, configure receiver as below:

 - notifier: `Slack`
 - to: a list of chat room names

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Slack
    to:
    - '#ops-alerts'
notifierSecretName: notifier-config
```


## Plivo
To receive SMS notifications via Plivo, create a Secret with the following keys:

| Name              | Description                             |
|-------------------|-----------------------------------------|
| PLIVO_AUTH_ID     | `Required` Plivo auth ID                |
| PLIVO_AUTH_TOKEN  | `Required` Plivo authentication token   |
| PLIVO_FROM        | `Required` Sender mobile number         |

```console
$ echo -n 'your-plivo-auth-id' > PLIVO_AUTH_ID
$ echo -n 'your-plivo-auth-token' > PLIVO_AUTH_TOKEN
$ echo -n 'your-plivo-from' > PLIVO_FROM
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./PLIVO_AUTH_ID \
    --from-file=./PLIVO_AUTH_TOKEN \
    --from-file=./PLIVO_FROM
secret "notifier-config" created
```
```yaml
apiVersion: v1
data:
  PLIVO_AUTH_ID: eW91ci1wbGl2by1hdXRoLWlk
  PLIVO_AUTH_TOKEN: eW91ci1wbGl2by1hdXRoLXRva2Vu
  PLIVO_FROM: eW91ci1wbGl2by1mcm9t
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T02:00:02Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "2606"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: f8dade1c-70dc-11e7-9b0b-080027503732
type: Opaque
```

Now, to receive notifications via Plivo, configure receiver as below:

 - notifier: `Plivo`
 - to: a list of receiver mobile numbers

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Plivo
    to:
    - +1-999-888-1234
notifierSecretName: notifier-config
```


## Pushover.net
To receive push notifications via Pushover.net, create a Secret with the following keys:

| Name               | Description                                                                    |
|--------------------|--------------------------------------------------------------------------------|
| PUSHOVER_TOKEN     | `Required` Pushover.net token.                                                 |
| PUSHOVER_USER_KEY  | `Required` User key or group key.                                              |
| PUSHOVER_TITLE     | `Optional` Message's title, otherwise your app's name is used.                 |
| PUSHOVER_URL       | `Optional` A supplementary URL to show with your message.                      |
| PUSHOVER_URL_TITLE | `Optional` A title for the supplementary URL, otherwise just the URL is shown. |
| PUSHOVER_PRIORITY  | `Optional` Send as -2 to generate no notification/alert, -1 to always send as a quiet notification, 1 to display as high-priority and bypass the user's quiet hours, or 2 to also require confirmation from the user.   |
| PUSHOVER_SOUND     | `Optional` The name of one of the sounds supported by device clients to override the user's default sound choice.   |


```console
$ echo -n 'your-pushover-token' > PUSHOVER_TOKEN
$ echo -n 'your-pushover-user-key' > PUSHOVER_USER_KEY
$ echo -n 'your-pushover-title' > PUSHOVER_TITLE
$ echo -n 'your-pushover-url' > PUSHOVER_URL
$ echo -n 'your-pushover-url-title' > PUSHOVER_URL_TITLE
$ echo -n 'your-pushover-priority' > PUSHOVER_PRIORITY
$ echo -n 'your-pushover-sound' > PUSHOVER_SOUND
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./PUSHOVER_TOKEN \
    --from-file=./PUSHOVER_USER_KEY \
    --from-file=./PUSHOVER_TITLE \
    --from-file=./PUSHOVER_URL \
    --from-file=./PUSHOVER_URL_TITLE \
    --from-file=./PUSHOVER_PRIORITY \
    --from-file=./PUSHOVER_SOUND
secret "notifier-config" created
```
```yaml
apiVersion: v1
data:
  PUSHOVER_PRIORITY: eW91ci1wdXNob3Zlci1wcmlvcml0eQ==
  PUSHOVER_SOUND: eW91ci1wdXNob3Zlci1zb3VuZA==
  PUSHOVER_TITLE: eW91ci1wdXNob3Zlci10aXRsZQ==
  PUSHOVER_TOKEN: eW91ci1wdXNob3Zlci10b2tlbg==
  PUSHOVER_URL: eW91ci1wdXNob3Zlci11cmw=
  PUSHOVER_URL_TITLE: eW91ci1wdXNob3Zlci11cmwtdGl0bGU=
  PUSHOVER_USER_KEY: eW91ci1wdXNob3Zlci11c2VyLWtleQ==
kind: Secret
metadata:
  creationTimestamp: 2017-08-04T05:13:07Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "33711872"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: 99df75a8-78d3-11e7-acfa-42010af00141
type: Opaque
```

Now, to receive notifications via Pushover.net, configure receiver as below:

 - notifier: `Pushover`
 - to: a list of devices where notifications will be sent. If list is empty, all devices will be notified.

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Pushover
    to:
    - my-phone
notifierSecretName: notifier-config
```


## Telegram
To receive notifications in your Telegram app, please follow the steps below:

- Sign up for [Telegram](https://telegram.org/).
- Now, use [@botfather](https://t.me/botfather) to create your own Telegram Bot and get the authorization token.
- Create a Kubernetes Secret with the following keys:

| Name               | Description                                                                    |
|--------------------|--------------------------------------------------------------------------------|
| TELEGRAM_TOKEN     | `Required` Telegram Bot authorization token.                                   |


```console
$ echo -n 'your-telegram-bot-token' > TELEGRAM_TOKEN
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./TELEGRAM_TOKEN
secret "notifier-config" created
```
```yaml
apiVersion: v1
data:
  TELEGRAM_TOKEN: NDkxoooooooooooooooooxJ
kind: Secret
metadata:
  creationTimestamp: 2018-01-16T13:41:04Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "768"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: e6066076-fac2-11e7-b3e7-0800276ee39b
type: Opaque
```

- Add the Bot as an Admin for your channel where notifications will be sent.

![telegram-bot](/docs/images/event-forwarder/telegram-bot.png)

Now, to receiver notifications in your **public** Telegram channels, configure receiver as below:

- notifier: `Telegram`
- to: a list of channels where notifications will be sent.

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Telegram
    to:
    - '@my-channel'
notifierSecretName: notifier-config
```

To receive notifications if your **private** channel, follow the steps below:

- Create a new public channel or change your private channel to public temporarily and assign a link.
- Now, run the following command to determine the `id` for your channel.

```console
TOKEN=<token>
curl -X POST https://api.telegram.org/bot${TOKEN}/sendMessage -d "chat_id=@channelusername&text=hello"

{"ok":true,"result":{"message_id":10,"chat":{"id":-1001210429328,"title":"mytest","username":"mytest123489","type":"channel"},"date":1516103121,"text":"hello"}}
```

- Make your channel private.

Now, to receiver notifications in your private Telegram channels, configure channle id as receiver like below:

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Telegram
    to:
    - '-1001210429328'
notifierSecretName: notifier-config
```


## Webhook Notifier
To receive chat notifications via Webhook, create a Secret with the following keys:

| Name                         | Description                                                                       |
|------------------------------|-----------------------------------------------------------------------------------|
| WEBHOOK_URL                  | `Required` URL of webhook server where notification is sent                       |
| WEBHOOK_USERNAME             | `Optional` Username for basic auth                                                |
| WEBHOOK_PASSWORD             | `Optional` Password for basic auth                                                |
| WEBHOOK_TOKEN                | `Optional` Token for bearer auth                                                  |
| WEBHOOK_CA_CERT_DATA         | `Optional` PEM encoded CA certificate used by Webhook server                      |
| WEBHOOK_CLIENT_CERT_DATA     | `Optional` PEM encoded client certificate used to authenticate to Webhook server  |
| WEBHOOK_CLIENT_KEY_DATA      | `Optional` PEM encoded client private key used to authenticate to Webhook server  |
| WEBHOOK_INSECURE_SKIP_VERIFY | `Optional` If set to `true`, skips SSL verification of Webhook server certificate |

```console
$ echo -n '' > WEBHOOK_URL
$ echo -n '' > WEBHOOK_TOKEN
$ echo -n '' > WEBHOOK_CA_CERT_DATA
$ kubectl create secret generic notifier-config -n kube-system \
    --from-file=./WEBHOOK_URL \
    --from-file=./WEBHOOK_TOKEN \
    --from-file=./WEBHOOK_CA_CERT_DATA
secret "notifier-config" created
```

```yaml
apiVersion: v1
data:
  WEBHOOK_URL: eW91ci1oaXBjaGF0LWF1dGgtdG9rZW4=
  WEBHOOK_TOKEN: eW91ci1oaXBjaGF0LWF1dGgtdG9rZW4=
  WEBHOOK_CA_CERT_DATA: eW91ci1oaXBjaGF0LWF1dGgtdG9rZW4=
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T01:54:37Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "2244"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: 372bc159-70dc-11e7-9b0b-080027503732
type: Opaque
```

Now, to receiver notifications via Webhook, configure receiver as below:

- notifier: `Webhook`
- to: a list of receiver names

```yaml
clusterName: unicorn
eventForwarder:
  receivers:
  - notifier: Webhook
    to:
    - ops-alerts
notifierSecretName: notifier-config
```


## Using multiple notifiers
Kubed supports using different notifiers in different scenarios. First add the credentials for the different notifiers in the same Secret `notifier-config` and deploy that to Kubernetes. Then in the Kubed cluster config, specify the appropriate notifier for each feature.

```yaml
apiVersion: v1
data:
  MAILGUN_API_KEY: eW91ci1tYWlsZ3VuLWFwaS1rZXk=
  MAILGUN_DOMAIN: eW91ci1tYWlsZ3VuLWRvbWFpbg==
  MAILGUN_FROM: bm8tcmVwbHlAZXhhbXBsZS5jb20=
  SLACK_AUTH_TOKEN: eW91ci1zbGFjay1hdXRoLXRva2Vu
kind: Secret
metadata:
  creationTimestamp: 2017-07-25T01:58:58Z
  name: notifier-config
  namespace: kube-system
  resourceVersion: "2534"
  selfLink: /api/v1/namespaces/kube-system/secrets/notifier-config
  uid: d2571817-70dc-11e7-9b0b-080027503732
type: Opaque


clusterName: unicorn
eventForwarder:
  warningEvents:
    handle: true
    namespaces:
    - kube-system
  receivers:
  - notifier: Mailgun
    to:
    - ops@example.com
  - notifier: Slack
    to:
    - '#ops-alerts'
notifierSecretName: notifier-config
```

## Next Steps
 - Want to keep an eye on your cluster with automated notifications? Setup Kubed [event forwarder](/docs/guides/cluster-events/event-forwarder.md).
 - Learn how to use Kubed to protect your Kubernetes cluster from disasters [here](/docs/guides/disaster-recovery/).
 - Need to keep configmaps/secrets synchronized across namespaces or clusters? Try [Kubed config syncer](/docs/guides/config-syncer/).
 - Out of disk space because of too much logs in Elasticsearch or metrics in InfluxDB? Configure [janitors](/docs/guides/janitors.md) to delete old data.
 - Wondering what features are coming next? Please visit [here](/docs/roadmap.md).
 - Want to hack on Kubed? Check our [contribution guidelines](/docs/CONTRIBUTING.md).
