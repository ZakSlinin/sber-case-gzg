# SBER-CASE-GZG 

### Service for calculating how much you spend on a subscription
#### Environment Configuration

Create your environment file from the example:

```bash
cp .env.example .env
```
Edit .env and configure your settings (database credentials, e-mail, etc.).

#### Run with Docker
```
$ docker-compose up
```

OR if you want the container to be in the background
```
$ docker-compose up -d 
```

After open ```localhost:8000/docs``` in your browser

**⚠️ Important Note**
 
Multiple starts may be required after the first build.  
Database migrations run automatically during startup. If the application fails to start on the first attempt, simply restart the container (`docker-compose up` again) until all migrations complete successfully.