import { Profile } from "./Profile"

export type User = {
    CreatedAt: Date;
    FirstName: string;
    LastName: string;
    Email: string;
    Profiles: Profile[];
}
