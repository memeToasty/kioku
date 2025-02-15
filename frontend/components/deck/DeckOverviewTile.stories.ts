import type { Meta, StoryObj } from "@storybook/react";
import DeckOverviewTile from "./DeckOverviewTile";

const meta: Meta<typeof DeckOverviewTile> = {
	title: "Deck/OverviewTile",
	component: DeckOverviewTile,
	tags: ["autodocs"],
	args: {
		id: "DeckId",
	},
};

export default meta;
type Story = StoryObj<typeof DeckOverviewTile>;

export const WithCardsDue: Story = {
	args: {
		name: "Example Deck",
		count: 2,
	},
};

export const NoCardsDue: Story = {
	args: {
		name: "Example Deck",
		count: 0,
	},
};
