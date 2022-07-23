
const StationView = ({station, setStation, lang}) => {


    return (
        <div style={{margin: '1em', border: '1px gray solid'}}>
            <button onClick={() => setStation(undefined)}
            style={{width: '100%', }}>close</button>
            <h3 style={{textAlign: 'center'}}>
                <span>{station.text[lang].name}</span></h3>

            <table>
                <tbody>
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

                </tbody>
            </table>

            <table>
                <thead>
                    <tr>
                        <th>x</th>
                        <th>y</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>{station.x}</td>
                        <td>{station.y}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}

export default StationView