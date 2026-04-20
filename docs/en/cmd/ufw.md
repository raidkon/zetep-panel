# `z-panel ufw`

Inspects **UFW** and prints guidance or template commands so your tunnel rules coexist with UFW.

Typical workflow: run the suggested checks, then apply the printed rules if they match your policy.

```bash
sudo z-panel ufw check
z-panel ufw help
```

Some operations need root to read firewall state accurately.
