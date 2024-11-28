# Songs API

## Installation:

Ensure that port 8080 is not allocated.

```bash
git clone https://github.com/aszxqaz/songs.git
cd songs
docker compose up --build
```

Navigate to [http://localhost:8080](http://localhost:8080)

## Song details service

To override the default mocked songs details service with a real implementation, set `SONG_INFO_SERVICE_SCHEME`, `SONG_INFO_SERVICE_HOST` and `SONG_INFO_SERVICE_PORT` environment variables
