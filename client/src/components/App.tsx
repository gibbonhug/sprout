import '../scss/react_components/App.scss';
import Flower from './Flower';

export const SERVER_ENDPOINT = 'http://localhost/3000';

function App() {
    return (
        <Flower
            colorPetal="aaaaaa"
            id={0}
            parentID={null}
            childID={null}
        ></Flower>
    );
}

export default App;
