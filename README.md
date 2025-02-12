# Opencast Capture Agent Display

Software backend for displays showing the current state of Opencast capture agents.

## How to build your Capture Agent Display

### List of components

- Case: https://github.com/virtUOS/opencast-ca-display/tree/3d-printed-case
- Raspberry Pi Compute Module 4: https://www.raspberrypi.com/products/compute-module-4/?variant=raspberry-pi-cm4001000
- Touchscreen: https://www.waveshare.com/7inch-hdmi-lcd-c.htm
- IO Board: https://www.waveshare.com/compute-module-4-poe-board-b.htm
- Miscellaneuos:
    - HDMI Cable: https://www.amazon.de/dp/B07R6CWPH1?th=1
    - USB Cable: https://amazon.de/dp/B095LS6S2Y
    - Ethernet Cable: https://www.amazon.de/ACT-Netzwerkkabel-gewinkelt-Flexibles-Datenzentren-Schwarz/dp/B0CYHG3HV2?th=1 (Alternatively, any other 90° right angled cable can be used)

### Instuctions

1. Put Raspberry Pi CM4 on the base board
2. Set the `boot` switch to `on`
3. Connect the board via USB-C cable
4. [Run `rpiboot`](https://github.com/raspberrypi/usbboot) to mount the CM4 file system
5. Start [Raspberry Pi Imager](https://www.raspberrypi.com/software/):
![Raspberry Pi Imager OS.png](assets/Raspberry%20Pi%20Imager%20OS.png)
- Device: Raspberry Pi 4
- OS: Raspberry Pi OS Lite (64 bit)
- Target: Select the CM4 Module Filesystem
- Additional settings:
    - Enable SSH public key
    - Set default SSH key
    - Disable telemetry
![Raspberry Pi settings.png](assets/Raspberry%20Pi%20settings.png)
6. Write image to CM 
7. Set `boot` to `off`
8. Boot Raspberry Pi
9. Run the Ansible playbook `prepare-os.yml`
10. Get MAC address (e.g. `ip a` on system)
11. Configure the `agent` setting in `/opt/opencast-ca-display/opencast-ca-display.yml` to the corresponding capture agent
12. Reboot the Raspberry Pi

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
