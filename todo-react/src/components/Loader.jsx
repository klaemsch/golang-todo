
// simple "Loading..." component with centered text
// is shown when the react waits for the backend to confirm the token
const Loader = () => {
  return (
    <section className="section">
      <div className="container">
        <div className="level">
          <div className="level-item has-text-centered">
            <p className="level">Loading...</p>
          </div>
        </div>
      </div>
    </section>
  )
}

export default Loader