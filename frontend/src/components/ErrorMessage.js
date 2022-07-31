const ErrorMessage = ({error}) => {
    const error_style = {
        color: 'white',
        backgroundColor: 'red',
        fontFamily: 'monospace',
        padding: '1em',
    }
    const message = error.msg.message

    return <div style={error_style}>
      Error: {message}
      <div>
      {
        message.startsWith('NetworkError') ? <div>The server may be offline.</div> : <></>
      }
      </div>
    </div>
}

export default ErrorMessage