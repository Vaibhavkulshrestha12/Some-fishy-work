import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [resources, setResources] = useState([]);
  const [error, setError] = useState(null);

  
  useEffect(() => {
    fetch('/api/resources')
      .then(response => {
        if (!response.ok) {
          throw new Error(`Error: ${response.status}`);
        }
        return response.json();
      })
      .then(data => setResources(data))
      .catch(err => setError(err.message));
  }, []);

  return (
    <div className="App">
      <h1>Resource Links</h1>
      {error && <p style={{ color: 'red' }}>Error fetching resources: {error}</p>}
      <div className="resources">
        {resources.map((resource, index) => (
          <div className="resource-card" key={index}>
            <img src={resource.icon} alt={`${resource.title} Icon`} />
            <h3>{resource.title}</h3>
            <p>{resource.description}</p>
            <a href={resource.url} target="_blank" rel="noopener noreferrer">Visit</a>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
