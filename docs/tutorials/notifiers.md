> New to Kubed? Please start [here](/docs/tutorials/README.md).

# Supported Notifiers
Kubed can send notifications via Email, SMS or Chat for various operations using [appscode/go-notify](https://github.com/appscode/go-notify) library. To connect to these services, you need to create a Secret with the appropriate keys. Then pass the secret name to Kubed by setting `notifierSecretName` field in Kubed cluster config.

## Hipchat
To receive chat notifications in Hipchat, create a Secret with the following key:

| Name                | Description                               |
|---------------------|-------------------------------------------|
| HIPCHAT_AUTH_TOKEN  | `Required` Hipchat [api access token](https://developer.atlassian.com/hipchat/guide/hipchat-rest-api/api-access-tokens). You can use room notification tokens, if you are planning to send notifications to a single room.   |

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

Now, to receiver notifications via Hipchat, configure receiver as below:
 - notifier: `hipchat`
 - to: a list of chat room names

```yaml
recycleBin:
  handleUpdates: false
  path: /tmp/kubed/trash
  receivers:
  - notifier: hipchat
    to:
    - ops-alerts
  ttl: 168h
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

Now, to receiver notifications via Mailgun, configure receiver as below:
 - notifier: `mailgun`
 - to: a list of email addresses

```yaml
recycleBin:
  handleUpdates: false
  path: /tmp/kubed/trash
  receivers:
  - notifier: mailgun
    to:
    - ops-alerts@example.com
  ttl: 168h
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

Now, to receiver notifications via SMTP, configure receiver as below:
 - notifier: `smtp`
 - to: a list of email addresses

```yaml
recycleBin:
  handleUpdates: false
  path: /tmp/kubed/trash
  receivers:
  - notifier: smtp
    to:
    - ops-alerts@example.com
  ttl: 168h
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

Now, to receiver notifications via SMTP, configure receiver as below:
 - notifier: `twilio`
 - to: a list of receiver mobile numbers

```yaml
recycleBin:
  handleUpdates: false
  path: /tmp/kubed/trash
  receivers:
  - notifier: twilio
    to:
    - +1-999-888-1234
  ttl: 168h
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

Now, to receiver notifications via Hipchat, configure receiver as below:
 - notifier: `slack`
 - to: a list of chat room names

```yaml
recycleBin:
  handleUpdates: false
  path: /tmp/kubed/trash
  receivers:
  - notifier: slack
    to:
    - '#ops-alerts'
  ttl: 168h
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

Now, to receiver notifications via SMTP, configure receiver as below:
 - notifier: `plivo`
 - to: a list of receiver mobile numbers

```yaml
recycleBin:
  handleUpdates: false
  path: /tmp/kubed/trash
  receivers:
  - notifier: plivo
    to:
    - +1-999-888-1234
  ttl: 168h
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


eventForwarder:
  warningEvents:
    handle: true
    namespaces:
    - kube-system
  receivers:
  - notifier: mailgun
    to:
    - ops@example.com
  - notifier: slack
    to:
    - #ops-alerts
recycleBin:
  handleUpdates: false
  path: /tmp/kubed/trash
  receivers:
  - notifier: mailgun
    to:
    - ops@example.com
  - notifier: slack
    to:
    - #ops-alerts
  ttl: 168h
notifierSecretName: notifier-config
```

## Next Steps
 - Learn how to use Kubed to take periodic snapshots of a Kubernetes cluster [here](/docs/tutorials/cluster-snapshot.md).
 - To setup a recycle bin for deleted and/or updated Kubernetes objects, please visit [here](/docs/tutorials/recycle-bin.md).
 - Need to keep some configuration synchronized across namespaces? Try [Kubed config syncer](/docs/tutorials/config-syncer.md).
 - Want to keep an eye on your cluster with automated notifications? Setup Kubed [event forwarder](/docs/tutorials/event-forwarder.md).
 - Out of disk space because of too much logs in Elasticsearch or metrics in InfluxDB? Configure [janitors](/docs/tutorials/janitors.md) to delete old data.
 - Wondering what features are coming next? Please visit [here](/ROADMAP.md).
 - Want to hack on Kubed? Check our [contribution guidelines](/CONTRIBUTING.md).
