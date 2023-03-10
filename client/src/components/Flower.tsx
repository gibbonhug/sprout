import { FlowerProps } from '../interfaces';

/**
 * A 'flower' (Flower-shaped SVG)
 * 
 * Styling is done inside Box scss since Flowers are dumb components.
 * 
 * @param props Props from JSON data. See FlowerProps documentation
 * @returns Colored SVG
 */
function Flower(props: FlowerProps) {
    /*
        Grab the prop to use as fill color for SVG
        (TypeScript expects style svg property to be an object, not a string)
    */
    const flowerStyle = {
        fill: props.colorPetal,
    };

    console.log("Flower " + props.id + "; color " + props.colorPetal);

    return (
        <svg
            width="100"
            height="100"
            viewBox="0 0 26.458333 26.458333"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
        >
            <g>
                <path
                    style={flowerStyle}
                    d="m 18.387133,17.846784 c 0,0 1.443415,7.002486 -0.970431,7.869576 -2.413847,0.86709 -5.472384,-5.555219 -5.472384,-5.555219 0,0 -5.1376801,4.795002 -7.0701481,3.064977 -1.932468,-1.730024 1.912184,-7.682597 1.912184,-7.682597 0,0 -6.58109692,-2.207483 -6.09971892,-4.804597 C 1.168013,8.1418103 8.0712029,8.6115473 8.0712029,8.6115473 c 0,0 -1.443417,-7.002484 0.970429,-7.86957398 2.4138461,-0.86709 5.4723831,5.55521798 5.4723831,5.55521798 0,0 5.13768,-4.795 7.070149,-3.064976 1.932467,1.730024 -1.912183,7.6825937 -1.912183,7.6825937 0,0 6.581095,2.207485 6.099717,4.804599 -0.481377,2.597114 -7.384565,2.127376 -7.384565,2.127376 z"
                />
            </g>
        </svg>
    );
}

export default Flower;