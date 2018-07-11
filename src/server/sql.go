package server

const insertASql = "INSERT IGNORE INTO partyA (name) VALUES (?)"

const selectAllASql = "SELECT id, name, (SELECT COUNT(*) FROM partyB WHERE partyAId=partyA.id) AS partyBNum FROM partyA LIMIT ?, ?"

const deleteASql = "DELETE FROM partyA WHERE id = ?"

const insertBSql = "INSERT IGNORE INTO partyB (name, partyAId, partyAUrl, partyBUrl) VALUES (?, ?, ?, ?)"

const selectBListSql = "SELECT id, name, partyAUrl, partyBUrl, clickCount FROM partyB WHERE partyAId = ? LIMIT ?, ?"

const deleteBSql = "DELETE FROM partyB WHERE id = ?"
