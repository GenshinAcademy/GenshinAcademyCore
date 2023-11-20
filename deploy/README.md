# Full deployment

It is recommended to deploy using ansible.

[Read documentation here.](ansible/README.md)

## Step-by-step manual guide

1. Create user `academy`.
2. Add user to groups:
    - sudo
    - www-data
3. Disable sudo password for user `academy`.
4. Add ssh keys to authorized keys for both `academy` and `root`.
5. Enable key access for ssh.
6. Disable password access for ssh.
7. Prepare folders for deployment:
    - /var/www/genshinacademy/prod/assets
    - /var/www/genshinacademy/prod/site
    - /var/www/genshinacademy/prod/server
    - /var/www/genshinacademy/dev/assets
    - /var/www/genshinacademy/dev/site
    - /var/www/genshinacademy/dev/server
8. Install Nginx
9. Configure nginx configurations
    - Templates
    - Certificates
10. Install Docker
11. Add user to groups:
    - docker
12. Login to Docker
13. Copy all required files
    - Docker compose configurations
    - .env files
14. Configure and enable services
15. Configure domain dns names (Currently using cloudflare)
