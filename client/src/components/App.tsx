import useSWR from 'swr';
import Flower from './Flower';
import { FlowerProps } from '../interfaces';

export const SERVER_ENDPOINT = 'http://localhost:3000';

const fetcher = (url: string) =>
    fetch(`${SERVER_ENDPOINT}/${url}`).then((r) => r.json());

function App() {
    // GET data from /flowers
    const { data, mutate } = useSWR<FlowerProps[]>('flowers', fetcher);

    console.log(data);
    return data?.map((flowerData) => {
        return <Flower key={flowerData.id} {...flowerData}></Flower>;
    });
}

export default App;
