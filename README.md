# Opencast Capture Agent Display

Software backend for displays showing the current state of Opencast capture agents.

## Build & Run

Make sure to install Go. This should be something like:

```
❯ dnf install golang
```

Then run the project with:

```
❯ go run main.go
```

…or build a static binary:

```
❯ go build
```

## Example

The display control in action:

https://github.com/virtUOS/opencast-ca-display/assets/1008395/ead22cd2-9d7a-4d26-97ae-e5744d23952d

- The display and laptop do not know about each other
- The laptop is running an Opencast capture agent
- When the laptop starts capturing video, the display shows an active recording

## Opencast User

To improve security, you can limit the access rights for the Opencast user by
creating a user which has only read access to the capture agent status API and
nothing else.

To do this, first create a new security rule in your Opencast's
`etc/security/mh_default_org.xml` allowing read access for a new role
`ROLE_CAPTURE_AGENT_CALENDAR`:

```xml
<!-- Enable capture agent updates and ingest -->
<sec:intercept-url pattern="/capture-admin/agents/**" method="GET" access="ROLE_ADMIN, ROLE_CAPTURE_AGENT, ROLE_CAPTURE_AGENT_CALENDAR" />
<sec:intercept-url pattern="/capture-admin/**" access="ROLE_ADMIN, ROLE_CAPTURE_AGENT" />
```

Next, go to the Opencast  REST Docs → `/user-utils` and fill out the form for
`POST /` with data like this:

- username: `ca-display`
- password: `secret-password`
- roles: `["ROLE_CAPTURE_AGENT_CALENDAR"]`

You should now be able to use this new user.
