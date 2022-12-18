import '../scss/react_components/Flower.scss';
import { FlowerProps } from '../interfaces';

/**
 * A 'flower' (currently a box)
 * @param props Props from JSON data. See FlowerProps documentation
 * @returns Colored box
 */
function Flower(props: FlowerProps) {

    // All styles except BG color written in SCSS file
    const dynamicStyle = {
        backgroundColor: "#" + props.colorPetal,
    }

    return <div className="flower" style={dynamicStyle}></div>
}

export default Flower;
