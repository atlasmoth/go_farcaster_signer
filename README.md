# Generate farcaster signers in go

## Setup

Create .env file and set the following variables

```bash
APP_FID="<0x116B52E794e1FC1DbceBBb2bc2C2ACffF6529b99>"
APP_MNEMONI<grain coin struggle satoshi gift rapid trash exit social govern cinnamon silly>"
```

Run start up command after installation

```bash
go run *.go
```

Make POST request to `/signer` to generate a signer

Make GET request to `/signer-status?token=$pollingToken` to poll status of the signer
