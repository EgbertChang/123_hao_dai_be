package server

const selectAInfo = "SELECT id, name FROM partyA WHERE id=5"

const selectBInfo = "SELECT id, name, partyAId, url FROM partyB WHERE id=6"

const selectBInfoList = "SELECT id, name, partyAId, url FROM partyB limit ?, ?"
