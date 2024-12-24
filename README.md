# Url shortening service build with Go:

Tech planned:

1. Go backend, built with Fiber framework. Deployed to AWS
2. Next.js simple UI, deployed on Vercel
3. Database: DynamoDB
   3.1 Update and delete logic
   3.2 QR code generation
   3.3 GDPR agreement before redirect
4. ? Cache: Redis (Cloud)
5. ? CDN or/and Load balancer between Client and API
6. ? Side-running Clean up service on DB
7. ? Analytics on shortening/retrieving operations
8. ? Visualization (ElasticSearch+Kibana)

In Progress:

1. Clean code

TODO (priority):

1. ~~DB deploy~~
2. ~~Zap log service connected~~
3. ~~Full API finish~~
4. ~~Project structure revamped~~
5. Go-routines and Channels best practices
6. API deployment on AWS
7. Redis cache work and it deployment
8. FE: Next.js simple UI and deployment
9. CDN
10. Simple Clean-up service
11. CDN or/and Load balancer between Client and API
