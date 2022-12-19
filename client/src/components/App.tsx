import useSWR from 'swr';
import Box from './Box';
import { BoxProps } from '../interfaces';

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
            {boxData &&
                boxData!.map((thisBox) => {
                    return <Box key={thisBox.id} {...thisBox}></Box>;
                })}
        </>
    );
}

export default App;
