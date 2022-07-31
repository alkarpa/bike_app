
const StationMap = ({ stations, filtered, active_low, active_high }) => {

    if (stations.length === 0) {
        return <></>
    }

    const min_x = stations.reduce((prev, cur) => Math.min(prev, cur.x), stations[0].x)
    const max_x = stations.reduce((prev, cur) => Math.max(prev, cur.x), stations[0].x)
    const diff_x = max_x - min_x
    //console.log('min_x', min_x, 'max_x', max_x,'diff', diff_x)
    const min_y = stations.reduce((prev, cur) => Math.min(prev, cur.y), stations[0].y)
    const max_y = stations.reduce((prev, cur) => Math.max(prev, cur.y), stations[0].y)
    const diff_y = max_y - min_y
    //console.log('min_y', min_y, 'max_y', max_y,'diff',diff_y)

    const scale = 90
    //console.log('diff_y/diff_x', diff_y/diff_x, 'y_scale', scale * (diff_y/diff_x))
    const adjusted_coords = stations.map(s => ({ id: s.id, x: (s.x - min_x) / diff_x * scale, y: (s.y - min_y) / diff_y * scale }))
    //console.log(adjusted_coords)

    const c_active = 'red'
    const c_filtered = 'gray'
    const c_rest = 'pink'

    const s_active = '4'
    const s_filtered = '3'
    const s_rest = '2'

    const color = (index) => (
        (index >= active_low && index < active_high) ? c_active : c_filtered
    )
    const size = (index) => (
        (index >= active_low && index < active_high) ? s_active : s_filtered
    )

    const filtered_coords = filtered.reduce((prev, cur, i) => ({ ...prev, [cur.id]: { id: cur.id, color: color(i), size: size(i) } }), {})
    const dots = adjusted_coords.map(c => {
        if (filtered_coords[c.id]) {
            return { ...c, ...filtered_coords[c.id] }
        }
        return { ...c, color: c_rest, size: s_rest }
    })


    return (
        <div>
            <div style={{ border: '1px black solid', 
            borderBottom: '0', 
            display: 'inline-block',
            width: `${scale}%`,
            textAlign: 'center',
            backgroundColor: 'gray'
            }}>Relative position map</div>
            <div style={{width: '100vw'}}>
                <svg style={{ width: `${scale}%`, aspectRatio: diff_x / diff_y, border: '1px brown solid' }}>
                    {
                        dots.map((c, i) => (
                            <rect key={`coord_${i}_${c.x}`} x={`${c.x + 5}%`} y={`${100 - (c.y + 5)}%`} width={c.size} height={c.size} fill={c.color} />
                        ))
                    }
                </svg>

            </div>
        </div>
    )

}

export default StationMap