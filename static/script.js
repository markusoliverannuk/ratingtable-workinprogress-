const playerName = document.querySelector('.selectedPlayerInfo h1');
        const ratingPoints = document.querySelector('.selectedPlayerInfo h2:nth-child(2)');
        const ratingPlace = document.querySelector('.selectedPlayerInfo h2:nth-child(3)');

        const tableRows = document.querySelectorAll('table tbody tr');

        tableRows.forEach(row => {
            const firstName = row.querySelector('td:nth-child(5)').textContent;

            row.querySelector('td:nth-child(5)').addEventListener('click', () => {
                
                const playerData = {
                    name: firstName, 
                    points: row.querySelector('td:nth-child(2)').textContent,
                    place: row.querySelector('td:nth-child(1)').textContent,
                };

                playerName.textContent = playerData.name;
                ratingPoints.textContent = `Reitingupunkte: ${playerData.points}`;
                ratingPlace.textContent = `Reitingu asetus: ${playerData.place}`;
            });
        });