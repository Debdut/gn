.. code:: console

    $ ./gn page create user/history
    $ ./gn api create user/profile 
    $ tree pages
    pages
    ├── api
    │   └── user
    │       └── profile.ts
    └── user
        └── history.tsx

    3 directories, 2 files

    $ cat pages/api/user/profile.ts

    import type { NextApiRequest, NextApiResponse } from 'next'

    type Profile = {
      prop: val
    }

    export default function handler(
      req: NextApiRequest,
      res: NextApiResponse<Profile>
    ) {
      res.status(200).json({ prop: 'value' })
    }

    $ cat pages/user/history.tsx

    import type { NextPage } from 'next'

    const History: NextPage = () => {
      return (
        <div className={styles.container}>
        </div>
      )
    }

    export default History

.. vim:ts=3 sts=3 sw=3 et ft=rst