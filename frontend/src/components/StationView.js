import { useEffect, useState } from "react"


const API_path = "http://localhost:8080"

const fetchStationStats = async (id) => {
    let response = await fetch(`${API_path}/station/${id}`)
    let json = await response.json()
    return json
}

const TopConnections = ({ identifier, connections, stationLang }) => (
    <div>
        <table style={{  display: 'inline-block', border: '2px solid lightblue', whiteSpace: 'nowrap' }}>
            <caption style={{ backgroundColor: 'lightblue', borderRadius: '5px 5px 0px 0px' }}>Top connections</caption>
            <tbody>
                {
                    connections.map((c, i) => (
                        <tr key={`connection_${i}_for_${identifier}`}>
                            <td>{i + 1}.</td>
                            <td>{stationLang[c.station_id]?.name}</td>
                            <td>{c.count}</td>
                        </tr>
                    ))
                }
            </tbody>
        </table>
    </div>
)
const StatView = ({ group, groupName, stationLang }) => {

    return (
        <div>
            <div style={{ backgroundColor: 'lightgray', textAlign: 'center' }}>{groupName}</div>
            <div style={{ display: 'flex', justifyContent: 'space-around', flexWrap: 'wrap' }}>
                <div>
                    <table style={{ whiteSpace: 'nowrap', display: 'inline-block' }}>
                        <tbody>
                            <tr>
                                <th>Total rides</th>
                                <td>{group.count}</td>
                            </tr>
                            <tr>
                                <th>Average distance</th>
                                <td><span style={{ fontFamily: 'monospace' }}> {group.average_distance} m</span></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <TopConnections identifier={groupName} connections={group.top_connections} stationLang={stationLang} />
            </div>
        </div>
    )
}

const StationsStats = ({ stats, stationLang }) => {
    const [filter, setFilter] = useState('all')
    if (!stats || stats === {} || Object.keys(stats).length === 0) {
        return <div>...</div>
    }
    const first = Object.keys(stats)[0]
    const groupings = Object.keys(stats[first])

    return (
        <div style={{ textAlign: 'center', backgroundColor: '#F6F6F6', flexGrow: 2 }}>
            <div>Filter statistics</div>
            {
                groupings.map(g => (
                    <button key={`grouping_${g}`} onClick={() => setFilter(g)}
                        style={filter === g ? { backgroundColor: 'white' } : {}}>{g}</button>
                ))
            }
            <div style={{ display: 'flex', justifyContent: 'space-around', flexWrap: 'wrap' }}>

                {
                    Object.keys(stats).map(key => (
                        <div key={`statistics_for_${key}`} style={{ margin: '1em', border: '2px solid gray', borderRadius: '2px' }}>
                            <div style={{ backgroundColor: 'gray', textAlign: 'center' }}>{key}</div>
                            {
                                Object.keys(stats[key]).filter(a => a === filter).map(group => (
                                    <StatView key={`statistics_for_${key}_${group}`}
                                        groupName={group}
                                        group={stats[key][group]}
                                        stationLang={stationLang} />
                                ))
                            }

                        </div>
                    ))
                }
            </div>
        </div>
    )
}

const StationView = ({ station, setStation, lang, stations }) => {
    const [stats, setStats] = useState({})

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetchStationStats(station.id)
            setStats(response)
        }
        fetchData()
    }, [station])

    const stationLang = stations.reduce((map, cur) => ({ ...map, [cur.id]: cur["text"][lang] }), {})
    return (
        <div style={{ margin: '1em', border: '1px gray solid' }}>
            <button onClick={() => setStation(undefined)}
                style={{ width: '100%', backgroundColor: 'darkred', color: 'white' }}>
                    back to station list
                </button>
            <h3 style={{ textAlign: 'center', backgroundColor: 'lightsteelblue', margin: 0 }}>
                <span>{station.text[lang].name}</span>
            </h3>
            <div style={{display: 'flex', flexWrap: 'wrap'}}>
                <div style={{flexGrow: 1}}>
                    <table style={{whiteSpace: 'nowrap', width: '100%'}}>
                        <tbody>
                            <tr>
                                <th>Address</th>
                                <td>{station.text[lang].address}</td>
                            </tr>
                            <tr>
                                <th>City</th>
                                <td>{station.text[lang].city || "N/A"}</td>
                            </tr>
                            <tr>
                                <th>Operator</th>
                                <td>{station.operator || "N/A"}</td>
                            </tr>
                            <tr>
                                <th>Capacity</th>
                                <td>{station.capacity}</td>
                            </tr>
                            <tr>
                                <th>x</th>
                                <td>{station.x}</td>
                            </tr>
                            <tr>
                                <th>y</th>
                                <td>{station.y}</td>
                            </tr>

                        </tbody>
                    </table>
                </div>

                <div>
                    <StationsStats stats={stats} stationLang={stationLang} />
                </div>
            </div>

        </div>
    )
}

export default StationView