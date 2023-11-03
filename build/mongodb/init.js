db = db.getSiblingDB('lizard');
db.createCollection('trends');

db.trends.createIndex({ short_url: 1 }, { unique: true })
db.trends.createIndex({ ai_id: 1 }, { unique: true })
db.trends.createIndex({ uid: 1 }, { unique: true })