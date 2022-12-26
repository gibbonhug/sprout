/**
 * See Go 'Flower' struct
 * - id: Unique unsigned ID
 * - colorPetal: Hex color of petals (with #)
*/
export interface FlowerProps {
    id: number;
    colorPetal: string;
}

/**
 * See Go 'Box' struct
 * - id: Unique unsigned ID
 * - flower: This box's flower, if it exists
 */
export interface BoxProps {
    id: number;
    flower: FlowerProps | null;
}

/**
 * -id: The same id as the dirt's parent Box
 */
export interface DirtProps {
    id: number;
}