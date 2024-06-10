#!/usr/bin/env node
const yargs = require("yargs");
const fs = require("fs");

let settings = null
const options = yargs
    .usage("Usage: -d <domain> -i <ip>")
    .option("a", { alias: "api", describe: "Your api key", type: "string", demandOption: false })
    .option("i", { alias: "ip", describe: "ip address for A record", type: "string", demandOption: false })
    .option("d", { alias: "domain", describe: "domain name for record", type: "string", demandOption: false })
    .option("f", { alias: "force", describe: "force override", type: "boolean", demandOption: false })
    .option("r", { alias: "remove", describe: "remove record", type: "string", demandOption: false })
    .argv;
    
let headers = {
    "Content-Type": "application/json"
}

const apiURL = (s) => {
    return `https://api.cloudflare.com/client/v4/${s}`
}

const recordLogic = async (full, domain, subdomain, ip) => {
    let zone = await findZoneByName(domain)
    if (zone == null) {
        console.log("no zone")
        return
    }
    let records = await getRecordsByZone(zone)
    let isExists = false
    let found = null
    if (records && records.result && records.result.map) {
        records.result.map((record) => {
            if (record.name == full) {
                if(record.type == "A") {
                    isExists = true
                    found = record
                    console.log(record)
                }
            }
        })
    }
    if(isExists) {
        console.log("**A record exists for the subdomain")
        if(options.force) {
            console.log("Forcing the dns")
            await deleteRecord(zone,found)
            await createRecord(zone, subdomain, ip)

        } else {
            console.log("**You can use -f parameter to force override dns record ")
        }
    } else {
        await createRecord(zone, subdomain, ip)
    }
}

const removeLogic = async (full, domain) => {
    let zone = await findZoneByName(domain)
    if (zone == null) {
        console.log("no zone")
        return
    }
    let records = await getRecordsByZone(zone)
    let isExists = false
    let found = null
    if (records && records.result && records.result.map) {
        records.result.map((record) => {
            if (record.name == full) {
                if(record.type == "A") {
                    isExists = true
                    found = record
                    console.log(record)
                }
            }
        })
    }
    if(isExists) {
        await deleteRecord(zone,found)
        console.log("**Record has been removed")
    } else {
        console.log("**Record doesnt exist")
    }
}


const findZoneByName = async (z) => {
    let res = await fetch(apiURL("zones"), { headers })
    let json = await res.json()
    let found = null
    if (res.status === 200) {
        json.result.map((zone) => {
            if (zone.name == z) {
                found = zone
            }
        })
    } else {
        console.log(res.status)
    }
    return found
}

const getRecordsByZone = async (z) => {
    let res = await fetch(apiURL(`zones/${z.id}/dns_records`), { headers })
    let json = await res.json()
    return json
}

const deleteRecord = async (zone, record) => {
    let resd = await fetch(apiURL(`zones/${zone.id}/dns_records/${record.id}`), {
        headers,
        method:"DELETE"
    })
    let json = await resd.json()
    return json
}



const createRecord = async (zone, domain, ip) => {
    let res = await fetch(apiURL(`zones/${zone.id}/dns_records`), {
        headers,
        method: "POST",
        body: JSON.stringify({
            type: "A",
            name: domain,
            content: ip,
            proxied: false,
            ttl: 1,
            comment: "created by api"
        })
    })
    let json = await res.json()
    console.log(json)
    return json
}

const startLogic = async () => {
    let isExists = fs.existsSync("./settings.json")
    if (!isExists) {
        fs.writeFileSync("./settings.json", JSON.stringify({
            api: ""
        }))
    }
    settings = JSON.parse(fs.readFileSync("./settings.json", "utf8"))

    if (options.api) {
        settings.api = options.api
        fs.writeFileSync("./settings.json", JSON.stringify(settings))
        console.log("Apikey has been updated")
        console.log("Use: flareup -d <domain> -i <ipaddress>")
        console.log("You can skip ip address parameter for using your own public ip")
        return
    }

    if (settings.api.length < 1) {
        console.log("Please set up your apikey first by using 'flareup -a <apiKey>'")
        return
    }
    headers["Authorization"] = `Bearer ${settings.api}`

    if(options.remove) {
        let targetURL = options.remove
        let url = targetURL.split(".")
        let subdomain = url[0]
        let domain = targetURL.replace(subdomain + ".", "")
        removeLogic(options.remove, domain)
        return
    }

    let ip = null
    if(options.ip) {
        ip = options.ip
    } else {
        let aip = await fetch("https://api.ipify.org?format=json")
        ip = (await aip.json()).ip
    }
    if(!options.domain) {
        console.log("Adding new record: flareup -d <sub.domain.com> -i <ipaddress>")
        console.log("[You can skip -i parameter to use your own public ip]")
        console.log("Deleting existing record: flareup -r <sub.domain.net>")
        console.log("Updating existing record by force: flareup -d <sub.domain.net> -i <ipaddress> -f")
        return
    }
    let targetURL = options.domain
    let url = targetURL.split(".")
    let subdomain = url[0]
    let domain = targetURL.replace(subdomain + ".", "")
    recordLogic(options.domain, domain, subdomain, ip)
}


startLogic()