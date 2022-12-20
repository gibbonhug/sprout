import useSWR from 'swr';
import Box from './Box';
import { BoxProps } from '../interfaces';
import '../scss/react_components/App.scss';

export const SERVER_ENDPOINT = 'http://localhost:3000';

const fetcher = (url: string) =>
    fetch(`${SERVER_ENDPOINT}/${url}`).then((r) => r.json());

function App() {
    // GET data from /boxes
    const { data: boxData, mutate } = useSWR<BoxProps[]>('boxes', fetcher);


    console.log(boxData);

    return (
        <>
            {!boxData && <div>Loading...</div>}
            <div id="box-grid">
                {boxData &&
                    boxData!.map((thisBox) => {
                        return <Box key={"box" + thisBox.id} {...thisBox}></Box>;
                    })}
            </div>
        </>
    );
}

export default App;
