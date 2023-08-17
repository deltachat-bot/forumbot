import { DateTime } from "luxon";

export namespace Types {
  export interface Community {
    name: string;
    description: string;
    posts: Post[];
  }

  export interface Post {
    id: string;
    title: string;
    createdAt: DateTime;
    body: string;
    comments: Comment[];
    commentCount: number;
    likes: number;
  }

  export interface Comment {
    id: string;
    body: string;
    comments: Comment[];
    createdAt: DateTime;
    likes: number;
  }
}

function mCommunity(
  name: string,
  description: string,
  posts: Types.Post[]
): Types.Community {
  return { name, description, posts };
}

function mPost(
  id: string,
  title: string,
  body: string,
  createdAt: DateTime,
  comments: Types.Comment[]
) {
  return {
    id,
    title,
    body,
    createdAt,
    comments,
    commentCount: countNestedComments(comments),
    likes: Math.floor(Math.random() * 10),
  };
}

export const MockData = [
  mCommunity("en", "community for the english speaking folks", [
    mPost(
      "455124254142",
      "webxdc is cool",
      "Webxdc is webapps inside of Messaging apps, it is a really cool technologie and I'm excited to try it out",
      DateTime.now().minus({ minutes: 10 }),
      [
        {
          id: "dfasafdr4w",
          body: "cool",
          comments: [
            {
              id: "323241412342",
              body: "nice",
              comments: [],
              createdAt: DateTime.now().minus({ minutes: 7 }),
              likes: Math.floor(Math.random() * 10),
            },
            {
              id: "ffsaddfs",
              body: "awesome!",
              comments: [],
              createdAt: DateTime.now().minus({ minutes: 7 }),
              likes: Math.floor(Math.random() * 10),
            },
          ],
          createdAt: DateTime.now().minus({ minutes: 8 }),
          likes: Math.floor(Math.random() * 10),
        },
        {
          id: "2354543523",
          body: "nice technology, I was using it recently for spectacular Lorem Ipsum....",
          comments: [
            {
              id: "099904",
              body: "+1",
              comments: [],
              createdAt: DateTime.now().minus({ minutes: 6 }),
              likes: Math.floor(Math.random() * 10),
            },
          ],
          createdAt: DateTime.now().minus({ minutes: 7 }),
          likes: Math.floor(Math.random() * 10),
        },
      ]
    ),
    mPost(
      "45678999",
      "deltachat the email chat",
      "Deltachat is like whatsapp and telegram but without the tracking an central control BS, just encrypted email and cool private chat experiences over webxdc technology",
      DateTime.now().minus({ minutes: 50 }),
      [
        {
          id: "834078312",
          body: "I tried it and it worked fine, super aswesome",
          comments: [
            {
              id: "0492902458",
              body: "can't say the same, my protonmail account won't work",
              comments: [
                {
                  id: "456396540",
                  body: "yeah you need to pay them for their bridge thingy that might work but I haven't tried it myself",
                  comments: [
                    {
                      id: "34578453819421",
                      body: "really? now I'm sad 😢",
                      comments: [],
                      createdAt: DateTime.now().minus({ minutes: 3 }),
                      likes: Math.floor(Math.random() * 10),
                    },
                  ],
                  createdAt: DateTime.now().minus({ minutes: 5 }),
                  likes: Math.floor(Math.random() * 10),
                },
              ],
              createdAt: DateTime.now().minus({ minutes: 17 }),
              likes: Math.floor(Math.random() * 10),
            },
          ],
          createdAt: DateTime.now().minus({ minutes: 18 }),
          likes: Math.floor(Math.random() * 10),
        },
        {
          id: "940971542789",
          body: "nice technology, I was using it recently for spectacular Lorem Ipsum....",
          comments: [],
          createdAt: DateTime.now().minus({ minutes: 27 }),
          likes: Math.floor(Math.random() * 10),
        },
        {
          id: "83218",
          body: "I use it with my family",
          comments: [],
          createdAt: DateTime.now().minus({ minutes: 16 }),
          likes: Math.floor(Math.random() * 10),
        },
      ]
    ),
  ]),
];
function countNestedComments(comments: Types.Comment[]): number {
  return (
    comments.reduce(
      (preValue, current) => preValue + countNestedComments(current.comments),
      0
    ) + comments.length
  );
}
