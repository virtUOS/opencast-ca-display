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