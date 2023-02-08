import React, { useState, useEffect } from 'react';
import axios from 'axios';


function MovieSearch() {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [selectedMovie, setSelectedMovie] = useState(null);

  const url = process.env.REACT_APP_API_URL;

  const handleSubmission = async (e) => {

    const updateSearchHistory = (searchParam) => {
      const history = JSON.parse(localStorage.getItem('searchHistory') || '[]');
      const isHistoryMaxed = history.length === 5;
      const workingHistory = isHistoryMaxed ? history.slice(1) : history;
      workingHistory.push(searchParam);
      localStorage.setItem('searchHistory', JSON.stringify(workingHistory));
    }

    const manageResults = (results) => {
      searchResults.splice(0, searchResults.length);
      if (Array.isArray(results)) {
        results.forEach((result) => {setSearchResults([...searchResults,{...result}])})
      } else {setSearchResults([...searchResults, {...results}])}};

    if (e !== undefined){e.preventDefault()};

    try {
      const response = await axios.get(`${url}${searchTerm}`);
      if (response.data.statusCode === 200) {
                const respData = response.data.data;
                manageResults(respData);
                updateSearchHistory(searchTerm);
            }
    } catch (error) {
      console.error(error);
    };
  }

  useEffect(() => {
    const fetchData = async () => {
      if (!searchTerm) {
        return;
      }
      await handleSubmission()
      
    };

    fetchData();
  }, [searchTerm]);

  const handleSearchInput = (event) => {
    setSearchTerm(event.target.value);
  };

  const handleMovieSelection = (title) => {
    setSelectedMovie(searchResults.find((movie) => movie.Title === title));
  };

  return (
    <div>
      <h2>Movie Search</h2>
      <form onSubmit={handleSubmission}>
        <input type="text" placeholder="Search by title" value={searchTerm} onChange={handleSearchInput}/>
        <button type="submit">Search</button>
      </form>
      <h3>Search Results</h3>
      <p>
        {searchResults.map((movie) => (
          <li key={movie.Title} onClick={() => handleMovieSelection(movie.Title)}>
            {movie.Title}
          </li>
        ))}
      </p>
      {selectedMovie && (
        <div>
          <h3>{selectedMovie.Title}</h3>
          <img src={selectedMovie.Poster} alt={selectedMovie.Title} />
          <p>{selectedMovie.Plot}</p>
          <p>AwardsWon : {selectedMovie.Awards}</p>
          <p>Genre : {selectedMovie.Genre}</p>
          <p>Director : {selectedMovie.Director}</p>
          <p>Actors : {selectedMovie.Actors}</p>
          <p>Writer : {selectedMovie.Writer}</p>
          <p>Year : {selectedMovie.Year}</p>
          <p>ReleaseDate : {selectedMovie.Released}</p>
          <p>Country : {selectedMovie.Country}</p>
          <p>Language : {selectedMovie.Language}</p>
          <p>Runtime : {selectedMovie.Runtime}</p>
          <p>Ratings : </p>
          <ul>
        {selectedMovie.Ratings.map((rating) => (
          <li><p>{rating.Source} : {rating.Value}</p>
          </li>
        ))}
      </ul>
        </div>
      )}
    </div>
  );
}

export default MovieSearch;
