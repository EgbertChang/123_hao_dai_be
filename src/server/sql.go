package server

const insertASql = "INSERT IGNORE INTO partyA (name) VALUES (?)"

const selectAllASql = "SELECT id, name, (SELECT COUNT(*) FROM partyB WHERE partyAId=partyA.id) AS partyBNum FROM partyA LIMIT ?, ?"

const deleteASql = "DELETE FROM partyA WHERE id = ?"

const selectAInfoSql = ""

const insertBSql = "INSERT IGNORE INTO partyB (name, partyAId, partyAUrl, partyBUrl) VALUES (?, ?, ?, ?)"

const selectBListSql = "SELECT id, name, partyAUrl, partyBUrl, clickCount FROM partyB WHERE partyAId = ? LIMIT ?, ?"

const deleteBSql = "DELETE FROM partyB WHERE id = ?"

const selectBInfoSql = ""

const insertProductSql = "INSERT IGNORE INTO product " +
	"(name, url, type, personalQualification, limitMin, limitMax, logoUrl, slogan," +
	" applyNumber, term, interest, lendingRate, credit, auditType, accountInType, applyStrategy) " +
	"VALUES " +
	"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

const deleteProductSql = "DELETE FROM product WHERE id = ?"

const updateProductSql = "UPDATE product Set name = ?, url = ?, type = ?, personalQualification = ?, " +
	"limitMin = ?, limitMax = ?, logoUrl = ?, slogan = ?, applyNumber = ?, term = ?, interest = ?, " +
	"lendingRate = ?, credit = ?, auditType = ?, accountInType = ?, applyStrategy = ? " +
	"WHERE id = ?"

const selectProductDetailSql = "SELECT * FROM product WHERE id = ?"

const selectProductListSql = "SELECT id, name, limitMin, limitMax, interest FROM product"
