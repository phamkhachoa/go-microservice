import { useState } from 'react';

export default function App() {
  const [apiResponse, setApiResponse] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchData = async () => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch('https://d2xjcju5cw0ag.cloudfront.net/api/inventory/v0/inventories/products/1');

      if (!response.ok) {
        throw new Error(`API responded with status ${response.status}`);
      }

      const data = await response.json();
      setApiResponse(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
      <div className="flex flex-col items-center p-6 max-w-2xl mx-auto">
        <h1 className="text-2xl font-bold mb-6">API Response Viewer</h1>

        <button
            onClick={fetchData}
            disabled={isLoading}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mb-6 disabled:bg-blue-300"
        >
          {isLoading ? 'Loading...' : 'Fetch Product Data'}
        </button>

        {error && (
            <div className="w-full bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              <p>Error: {error}</p>
            </div>
        )}

        {apiResponse && (
            <div className="w-full">
              <h2 className="text-xl mb-2">API Response:</h2>
              <pre className="bg-gray-100 p-4 rounded overflow-auto max-h-96 w-full">
            {JSON.stringify(apiResponse, null, 2)}
          </pre>
            </div>
        )}
      </div>
  );
}