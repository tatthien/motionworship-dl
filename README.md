# MotionWorship Downloader

A simple tool that helps you download multiple HD video at the same time from [MotionWorship](https://motionworship.com/) website easily.

## Installation

```bash
go install github.com/tatthien/motionworship-dl
```

## Usage

1. Before downloading video, you need to login to MotionWorship website and get the authentication cookie from that site.
2. Export `COOKIE` environment variable.
3. Run `motionworship-dl <motionworship-url>`. E.g: `motionworship-dl https://www.motionworship.com/vhs-glitch-collection/`

