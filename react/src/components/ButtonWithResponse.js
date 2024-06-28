import React, { useState } from 'react';
import axios from 'axios';

const ButtonWithResponse = ({ baseUrl, paths }) => {
  const [responses, setResponses] = useState({});
  const [errors, setErrors] = useState({});

  const handleClick = async (path) => {
    try {
      const res = await axios.get(`${baseUrl}${path}`);
      setResponses((prev) => ({ ...prev, [path]: res }));
      setErrors((prev) => ({ ...prev, [path]: null }));
    } catch (err) {
      setErrors((prev) => ({ ...prev, [path]: err.response }));
      setResponses((prev) => ({ ...prev, [path]: null }));
    }
  };

  const getBoxStyle = (path) => {
    const response = responses[path];
    const error = errors[path];
    if (response) {
      if (response.status >= 200 && response.status < 300) {
        return { backgroundColor: 'rgba(0, 255, 0, 0.1)', color: 'black' };
      }
    } else if (error) {
      return { backgroundColor: 'rgba(255, 0, 0, 0.1)', color: 'red' };
    }
    return { backgroundColor: 'lightgray', color: 'black' };
  };

  return (
    <div>
      {paths.map((path) => (
        <div key={path}>
          <button onClick={() => handleClick(path)}>{`${baseUrl}${path}`}</button>
          <div style={{ ...getBoxStyle(path), padding: '10px', marginTop: '10px' }}>
            {responses[path] && (
              <div>
                <p>Status: {responses[path].status}</p>
                <p>Data: {JSON.stringify(responses[path].data)}</p>
              </div>
            )}
            {errors[path] && (
              <div>
                <p>Status: {errors[path].status}</p>
                <p>Error: {JSON.stringify(errors[path].data)}</p>
              </div>
            )}
          </div>
        </div>
      ))}
    </div>
  );
};

export default ButtonWithResponse;
