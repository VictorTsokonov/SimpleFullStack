import React, { useState, useEffect } from 'react';
import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080"; // Default to localhost for local dev


const App = () => {
  const [games, setGames] = useState([]);
  const [name, setName] = useState('');
  const [releaseDate, setReleaseDate] = useState('');

  // Fetch games from the backend
  const fetchGames = async () => {
    try {
      let response = await axios.get(`${API_URL}/game`);
      console.log(response)
      if(response?.data != null)setGames(response.data);
      if(response == null)setGames([])

    } catch (error) {
      console.error('Error fetching games:', error);
    }
  };

  // Add a new game
  const addGame = async () => {
    if (!name || !releaseDate) return;

    try {
      await axios.post(`${API_URL}/game`, {
        name,
        release_date: parseInt(releaseDate, 10),
      });
      setName('');
      setReleaseDate('');
      fetchGames();
    } catch (error) {
      console.error('Error adding game:', error);
    }
  };


  const deleteGame = async (id) => {
    try {
      await axios.delete(`${API_URL}/game?id=${id}`);

      setGames((prevGames) => prevGames.filter((game) => game.id !== id));

    } catch (error) {
      console.error('Error deleting game:', error);
    }
  };


  useEffect(() => {
    fetchGames();
  }, []);

  return (
    <div style={{ padding: '20px', fontFamily: 'Arial, sans-serif' }}>
      <h1>Game Library</h1>

      <div style={{ marginBottom: '20px' }}>
        <h2>Add a New Game</h2>
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          style={{ marginRight: '10px', padding: '5px' }}
        />
        <input
          type="text"
          placeholder="Release Date"
          value={releaseDate}
          onChange={(e) => setReleaseDate(e.target.value)}
          style={{ marginRight: '10px', padding: '5px' }}
        />
        <button onClick={addGame} style={{ padding: '5px 10px' }}>
          CREATE
        </button>
      </div>

      <div>
        <h2>List of Games</h2>
        {games.map((game) => (
          <div key={game.id} style={{ marginBottom: '10px', border: '1px solid #ccc', padding: '10px', borderRadius: '5px' }}>
            <p><strong>GAME:</strong> {game.name}</p>
            <p><strong>Release Date:</strong> {game.release_date}</p>
            <button onClick={() => deleteGame(game.id)} style={{ padding: '5px 10px', backgroundColor: '#f00', color: '#fff', border: 'none', borderRadius: '3px' }}>
              DELETE
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default App;