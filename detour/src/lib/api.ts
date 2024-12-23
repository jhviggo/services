import { addError, Levels } from './errors';

interface Blog {
  slug: string;
  headline: string;
  image: string;
  date: Date;
  text: string;
}

export interface Timeline {
  id: number;
  itemId: string;
  date: string;
  location: string;
  city: string;
  image: string;
  content: string;
}

const SERVICES_API_URL = '';

export async function getTimelines(): Promise<Timeline[]> {
  try {
    const res = await fetch('/timeline.json')
    return await res.json();
  } catch (e) {
    addError({
      level: Levels.ERROR,
      message: 'Unable to get timeline'
    });
  }
  return [];
}

export async function getContentFile(fileName: string): Promise<string> {
  try {
    const res = await fetch(`/contents/${fileName}`);
    return res.text();
  } catch (e) {
    addError({
      level: Levels.ERROR,
      message: 'Unable to get timeline'
    });
  }
  return '';
}

export async function registerVisitor(visitorId: string) {
  try {
    const res = await fetch(`${SERVICES_API_URL}/register-visitor`, {
      method: 'POST',
      body: JSON.stringify({
        visitor: visitorId,
      }),
    });
  } catch (e) {
    console.warn('failed to register visitor');
  }
}
