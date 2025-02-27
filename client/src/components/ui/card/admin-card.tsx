import * as React from 'react';
import Link from "next/link"

import { Card, CardContent } from '@/components/ui/card';

type Props = {
  cardTitle: string,
  cardTitleRuby: string,
  cardDescription: string,
  cardLink:  string
};

const AdminCard: React.FC<Props> = (props) => {
  return (
    <Link href={props.cardLink} passHref>
      <Card className="w-full max-w-md mx-auto bg-white dark:bg-gray-800 overflow-hidden">
        <CardContent className="p-6">
          <div className="flex items-center mb-4">
            <div className="relative w-20 h-20 mr-4 bg-gray-100 dark:bg-gray-700 rounded-lg overflow-hidden flex-shrink-0">
              <div className="absolute inset-0 flex items-end justify-center">
                <div className="w-full h-3/4 relative">
                  {/* Mountain */}
                  <div className="absolute bottom-0 left-0 w-full h-3/4 bg-gray-300 dark:bg-gray-600" style={{ clipPath: 'polygon(15% 100%, 50% 0%, 85% 100%)' }}></div>
                  {/* Torii */}
                  <div className="absolute bottom-0 left-1/2 -translate-x-1/2 w-3/5 h-3/5">
                    <div className="absolute top-0 left-0 w-full h-1 bg-red-700 dark:bg-gray-200"></div>
                  </div>
                  <div className="absolute bottom-0 left-1/2 -translate-x-1/2 w-7/12 h-3/5">
                    <div className="absolute top-2 left-0 w-full h-1 bg-red-700 dark:bg-gray-200"></div>
                  </div>
                  <div className="absolute bottom-0 left-1/2 -translate-x-1/2 w-1/2 h-3/5">
                    <div className="absolute top-0 left-0 w-1 h-full bg-red-700 dark:bg-gray-200"></div>
                    <div className="absolute top-0 right-0 w-1 h-full bg-red-700 dark:bg-gray-200"></div>
                  </div>
                </div>
              </div>
            </div>
            <ruby className="text-xl font-semibold text-gray-800 dark:text-gray-200">
              {props.cardTitle}
              <rt>{props.cardTitleRuby}</rt>
            </ruby>
          </div>
          <p className="text-sm text-gray-600 dark:text-gray-400 mb-4">{props.cardDescription}</p>
        </CardContent>
      </Card>
    </Link>
  );
};

export {AdminCard};