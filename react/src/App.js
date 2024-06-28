import React from 'react';
import ButtonWithResponse from './components/ButtonWithResponse';

const services = [
  { name: 'Java', baseUrl: '/java' },
  { name: 'Golang', baseUrl: '/golang' },
  { name: 'NodeJS', baseUrl: '/nodejs' },
  { name: 'Python', baseUrl: '/python' },
  { name: 'DotNet', baseUrl: '/dotnet' },
  { name: 'PHP', baseUrl: '/php' },
];

const paths = ['/', '/sleep/3', '/status/random', '/exception', '/users'];

const App = () => {
  return (
    <div>
      {services.map((service) => (
        <div key={service.name}>
          <h2>{service.name} Service</h2>
          <ButtonWithResponse baseUrl={service.baseUrl} paths={paths} />
        </div>
      ))}
    </div>
  );
};

export default App;
