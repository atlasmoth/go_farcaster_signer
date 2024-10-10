# Generate farcaster signers in go

## Setup

Create .env file and set the following variables

```bash
APP_FID="<DEVELOPER'S FID>"
APP_MNEMONIC="<DEVELOPER'S MNEMONIC>"
```

Run start up command after installation

```bash
go run *.go
```

Make POST request to `/signer` to generate a signer
Make GET request to `/signer-status?token=$pollingToken` to poll status of the signer
