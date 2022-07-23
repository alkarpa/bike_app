
const StationMap = ({stations, active_low, active_high}) => {

    if (stations.length === 0) {
        return <></>
    }

    const min_x = stations.reduce( (prev, cur) => Math.min(prev,  cur.x), stations[0].x )
    const max_x = stations.reduce( (prev, cur) => Math.max(prev,  cur.x), stations[0].x)
    const diff_x = max_x-min_x
    //console.log('min_x', min_x, 'max_x', max_x,'diff', diff_x)
    const min_y = stations.reduce( (prev, cur) => Math.min(prev,  cur.y), stations[0].y )
    const max_y = stations.reduce( (prev, cur) => Math.max(prev,  cur.y), stations[0].y)
    const diff_y = max_y-min_y
    //console.log('min_y', min_y, 'max_y', max_y,'diff',diff_y)

    const scale = 90
    //console.log('diff_y/diff_x', diff_y/diff_x, 'y_scale', scale * (diff_y/diff_x))
    const adjusted_coords = stations.map( s => ({ x: (s.x - min_x) / diff_x * scale, y: (s.y - min_y) / diff_y * scale }) )
    //console.log(adjusted_coords)

    const color = (index) => (
        (index >= active_low && index < active_high) ? 'red' : 'gray'
    )

    return (
        <svg style={{width: `${scale}%`, aspectRatio: diff_x/diff_y, border: '1px brown solid'}}>
                {
                    adjusted_coords.map( (c,i) => (
                        <rect key={`coord_${i}_${c.x}`} x={`${c.x+5}%`} y={`${100-(c.y+5)}%`} width={3} height={3} fill={color(i)} />
                    ) )
                }
            </svg>
    )

}

export default StationMap