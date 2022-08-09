const SVGWheel = ({ radius }) => {
    const center = radius
    return (
        <g className='Rotating' style={{ transformOrigin: `${center}px ${center}px` }}>
            <circle cx={center} cy={center} r={radius} fill="transparent" />
            <line x1={center} y1={center - radius} x2={center} y2={center + radius} />
            <line x1={center - radius} y1={center} x2={center + radius} y2={center} />
        </g>
    )
}

const SVGBike = () => {
    const wheel_radius = 50
    const back_wheel_center = { x: 60, y: 140 }
    const front_wheel_center = { x: back_wheel_center.x + wheel_radius * 3.3, y: back_wheel_center.y }
    const pedal = { x: back_wheel_center.x + wheel_radius * 1.4, y: back_wheel_center.y + wheel_radius * 0.2 }
    const handle_joint = { x: front_wheel_center.x - wheel_radius * 0.5, y: back_wheel_center.y - wheel_radius * 1.5 }
    const seat_joint = { x: pedal.x - wheel_radius * 0.3, y: handle_joint.y + wheel_radius * 0.3 }
    const seat = { x: seat_joint.x - wheel_radius * 0.1, y: seat_joint.y - wheel_radius * 0.3 }
    const handle = { x: handle_joint.x - wheel_radius * 0.1, y: handle_joint.y - wheel_radius * 0.3 }

    const lines = [
        { from: back_wheel_center, to: pedal },
        { from: front_wheel_center, to: handle_joint },
        { from: pedal, to: handle_joint },
        { from: pedal, to: seat_joint },
        { from: handle_joint, to: seat_joint },
        { from: back_wheel_center, to: seat_joint },
        { from: seat_joint, to: seat },
        { from: handle_joint, to: handle }
    ]

    const Line = ({ line }) => (
        <line x1={line.from.x} y1={line.from.y} x2={line.to.x} y2={line.to.y} />
    )

    const Seat = () => (
        <line x1={seat.x - wheel_radius * 0.4} y1={seat.y} x2={seat.x + wheel_radius * 0.4} y2={seat.y} />
    )
    const Handle = () => (
        <g>
            <line x1={handle.x + wheel_radius * 0.2} y1={handle.y - wheel_radius * 0.2} x2={handle.x} y2={handle.y} />
            <line x1={handle.x + wheel_radius * 0.2} y1={handle.y - wheel_radius * 0.2} x2={handle.x - wheel_radius * 0.2} y2={handle.y - wheel_radius * 0.25} />
        </g>
    )

    return (
        <svg style={{ position: 'relative', width: '300px', height: '250px', stroke: 'black', strokeWidth: '10px' }}>

            {
                lines.map((l,i) => (
                    <Line key={`loading_bike_line_${i}`} line={l} />
                ))
            }

            <Seat />
            <Handle />

            <g style={{ transform: `translate(${back_wheel_center.x - wheel_radius}px, ${back_wheel_center.y - wheel_radius}px)` }}>
                <SVGWheel radius={wheel_radius} />
            </g>
            <g style={{ transform: `translate(${front_wheel_center.x - wheel_radius}px, ${front_wheel_center.y - wheel_radius}px)` }}>
                <SVGWheel radius={wheel_radius} />
            </g>

        </svg>
    )
}

const LoadingPlaceholder = () => (
    <div style={{ display: 'grid', gridTemplateRows: 'min-content 1fr', justifyItems: 'center', marginTop: '100px' }}>
        <h1 style={{ textAlign: 'center', margin: '0' }}>Loading</h1>
        <SVGBike />
    </div>
)

export default LoadingPlaceholder