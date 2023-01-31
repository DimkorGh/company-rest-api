db.createCollection('company');
db.company.createIndex({name: 1}, {unique: true})
