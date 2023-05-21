## Info

Daemons collection for different stages of this project.

- `ga_server` - Server Executable version.
- `ga_server_docker` - Server Docker version. (Current)

```bash
# Reload daemons after creating configs in /etc/systemd/system
sudo systemctl daemon-reload

# Enable service for run at boot
sudo systemctl enable ${SERVICE_NAME}

# Start service
sudo service ${SERVICE_NAME} start

# Check service status
sudo service ${SERVICE_NAME} status
```