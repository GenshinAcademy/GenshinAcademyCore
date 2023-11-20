# Deploy with ansible

1. Prepare all files in [playbooks/files](playbooks/files)
    - SSL certificates for nginx
    - SSH public keys
2. Copy [playbooks/full-deployment.dist.yml](playbooks/full-deployment.dist.yml) to `playbooks/full-deployment.yml` to
   not accidentally expose secret variables.

```bash  
cp playbooks/full-deployment.dist.yml playbooks/full-deployment.yml
```

3. Prepare all commented variables in `playbooks/full-deployment.yml`
4. Copy [hosts.txt.dist](hosts.txt.dist) to `hosts.txt` to not accidentally expose secret variables and prepare host
   configuration in `hosts.txt`.

```bash
cp hosts.txt.dist hosts.txt 
```

5. Run ansible playbook with following commands:

```bash
ansible-playbook -i hosts.txt playbooks/full-deployment.yml
```
