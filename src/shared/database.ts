import { Database } from 'sqlite3';

export let dbInstance: Database;

export async function initializeDatabase() {
  const databaseLocation: string = process.env.DATABASE_PATH!;

  if (!databaseLocation) {
    throw new Error('Unable to conenct to database.');
  }

  dbInstance = new Database(databaseLocation, (error) => {
    if (error) {
      throw new Error(error.message);
    }
  });
}
